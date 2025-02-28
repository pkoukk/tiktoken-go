package tiktoken

import (
	"reflect"
	"unsafe"
)

func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func stringToRunes(s string) (b []rune) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}

func runesToString(b []rune) string {
	return *(*string)(unsafe.Pointer(&b))
}
