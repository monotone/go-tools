package cmm

import "os"
import "path"

func FileExist(filepath string) bool {
	stat, err := os.Lstat(filepath)
	if err == nil && stat != nil {
		return true
	}
	return false
}

func FileTittle(filepath string) string {
	_, filename := path.Split(filepath)
	return filename
}
