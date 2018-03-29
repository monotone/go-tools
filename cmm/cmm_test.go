package cmm

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"unsafe"
)

func TestUnZipTar(t *testing.T) {
	os.RemoveAll("./test")
	fmt.Println(UnZipTar("./testdata/test.tar.gz", "./"))
}

func CArrayToGoSlice(begin uintptr, size int) []byte {
	var theGoSlice []byte
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&theGoSlice)))
	sliceHeader.Cap = size
	sliceHeader.Len = size
	sliceHeader.Data = begin
	return theGoSlice
}

func TestGetPhysicsInterfaces(t *testing.T) {
	fmt.Println(GetPhysicsInterfaces(true))
}
