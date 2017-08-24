package cmm

import "testing"
import "fmt"

func TestUnZipTar(t *testing.T) {
	fmt.Println(UnZipTar("./test.tar.gz", "./"))
}
