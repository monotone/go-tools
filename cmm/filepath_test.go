package cmm

import "testing"
import "fmt"

func TestFileTittle(t *testing.T) {
	fmt.Println(FileTittle("http://fffff/test.tar.gz"))

	fmt.Println(FileTittle("abc/test.tar.gz"))

	fmt.Println(FileTittle("../test.tar.gz"))
}
