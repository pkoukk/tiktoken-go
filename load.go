package tiktoken

import (
	"bytes"
	"crypto/sha1"
	_ "embed"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

var (
	//go:embed cl100k.gob
	cl100k []byte
	//go:embed p50k.gob
	p50k          []byte
	embedded_maps = func() (s struct {
		Cl100k_base map[string]int
		P50k_base   map[string]int
	}) {
		dec := gob.NewDecoder(bytes.NewReader(cl100k))
		if err := dec.Decode(&s.Cl100k_base); err != nil {
			panic(err)
		}
		dec = gob.NewDecoder(bytes.NewReader(p50k))
		if err := dec.Decode(&s.P50k_base); err != nil {
			panic(err)
		}
		return
	}()
)

type BpeLoader interface {
	LoadTiktokenBpe(tiktokenBpeFile string) (map[string]int, error)
}

func readFile(blobpath string) ([]byte, error) {
	if !strings.HasPrefix(blobpath, "http://") && !strings.HasPrefix(blobpath, "https://") {
		file, err := os.Open(blobpath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		return ioutil.ReadAll(file)
	}
	// avoiding blobfile for public files helps avoid auth issues, like MFA prompts
	resp, err := http.Get(blobpath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func readFileCached(blobpath string) ([]byte, error) {
	var cacheDir string
	if os.Getenv("TIKTOKEN_CACHE_DIR") != "" {
		cacheDir = os.Getenv("TIKTOKEN_CACHE_DIR")
	} else if os.Getenv("DATA_GYM_CACHE_DIR") != "" {
		cacheDir = os.Getenv("DATA_GYM_CACHE_DIR")
	} else {
		cacheDir = filepath.Join(os.TempDir(), "data-gym-cache")
	}

	if cacheDir == "" {
		// disable caching
		return readFile(blobpath)
	}

	cacheKey := fmt.Sprintf("%x", sha1.Sum([]byte(blobpath)))

	cachePath := filepath.Join(cacheDir, cacheKey)
	if _, err := os.Stat(cachePath); err == nil {
		return ioutil.ReadFile(cachePath)
	}

	contents, err := readFile(blobpath)
	if err != nil {
		return nil, err
	}

	os.MkdirAll(cacheDir, os.ModePerm)
	tmpFilename := cachePath + "." + uuid.New().String() + ".tmp"
	if err := ioutil.WriteFile(tmpFilename, contents, os.ModePerm); err != nil {
		return nil, err
	}
	return contents, os.Rename(tmpFilename, cachePath)
}

func loadTiktokenBpe(tiktokenBpeFile string) (map[string]int, error) {
	contents, err := readFileCached(tiktokenBpeFile)
	if err != nil {
		return nil, err
	}

	bpeRanks := make(map[string]int)
	for _, line := range strings.Split(string(contents), "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		token, err := base64.StdEncoding.DecodeString(parts[0])
		if err != nil {
			return nil, err
		}
		rank, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		bpeRanks[string(token)] = rank
	}
	return bpeRanks, nil
}

type defaultBpeLoader struct{}

func (l *defaultBpeLoader) LoadTiktokenBpe(tiktokenBpeFile string) (map[string]int, error) {
	switch tiktokenBpeFile {
	case "https://openaipublic.blob.core.windows.net/encodings/cl100k_base.tiktoken":
		return embedded_maps.Cl100k_base, nil
	case "https://openaipublic.blob.core.windows.net/encodings/p50k_base.tiktoken":
		return embedded_maps.P50k_base, nil
	default:
		return loadTiktokenBpe(tiktokenBpeFile)
	}
}

func NewDefaultBpeLoader() BpeLoader {
	return &defaultBpeLoader{}
}
