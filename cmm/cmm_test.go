package cmm

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestUnZipTar(t *testing.T) {
	fmt.Println(UnZipTar("./test.tar.gz", "./"))
}

func CArrayToGoSlice(begin uintptr, size int) []byte {
	var theGoSlice []byte
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&theGoSlice)))
	sliceHeader.Cap = size
	sliceHeader.Len = size
	sliceHeader.Data = begin
	return theGoSlice
}
