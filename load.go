package tiktoken

import (
    "crypto/sha1"
    "crypto/sha256"
    "encoding/base64"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "strings"

    "github.com/google/uuid"
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

    // Optional integrity check against known hashes (opt-in via env)
    if os.Getenv("TIKTOKEN_VERIFY_HASH") == "1" {
        if expected, ok := expectedBPEHashes[tiktokenBpeFile]; ok && expected != "" {
            actual := fmt.Sprintf("%x", sha256.Sum256(contents))
            if actual != expected {
                return nil, fmt.Errorf("bpe file hash mismatch for %s: expected %s got %s", tiktokenBpeFile, expected, actual)
            }
        }
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
	return loadTiktokenBpe(tiktokenBpeFile)
}

func NewDefaultBpeLoader() BpeLoader {
    return &defaultBpeLoader{}
}

// expectedBPEHashes are the upstream SHA-256 digests for official BPE files
var expectedBPEHashes = map[string]string{
    "https://openaipublic.blob.core.windows.net/encodings/r50k_base.tiktoken": "306cd27f03c1a714eca7108e03d66b7dc042abe8c258b44c199a7ed9838dd930",
    "https://openaipublic.blob.core.windows.net/encodings/p50k_base.tiktoken": "94b5ca7dff4d00767bc256fdd1b27e5b17361d7b8a5f968547f9f23eb70d2069",
    "https://openaipublic.blob.core.windows.net/encodings/cl100k_base.tiktoken": "223921b76ee99bde995b7ff738513eef100fb51d18c93597a113bcffe865b2a7",
    "https://openaipublic.blob.core.windows.net/encodings/o200k_base.tiktoken": "446a9538cb6c348e3516120d7c08b09f57c36495e2acfffe59a5bf8b0cfb1a2d",
}
