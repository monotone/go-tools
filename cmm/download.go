package cmm

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/pkg/errors"
)

// DownloadToFile 下载内容到文件，要求传入的文件必须不存在
func DownloadToFile(urlStr, filename string) error {
	var err error

	// 下载
	resp, err := http.Get(urlStr)
	if err != nil {
		return errors.Wrap(err, "request failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("resp status code not 200 : " + resp.Status)
	}

	// 准备好目标文件
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_EXCL|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	defer func() {
		if err != nil {
			os.Remove(filename)
		}
	}()

	// 写入文件
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return errors.Wrap(err, "write to file failed")
	}

	return nil
}

// Download 下载指定urlstr的内容到当前目录下，文件名从urlStr末尾获取
func Download(urlStr string) (string, error) {
	_, filename := path.Split(urlStr)
	filename = "./" + filename
	return filename, DownloadToFile(urlStr, filename)
}

// DownloadToDir 下载指定urlStr的内容到指定目录下，文件名从urlStr末尾获取
func DownloadToDir(urlStr, dir string) (string, error) {
	_, filename := path.Split(urlStr)
	filename = path.Join(dir, filename)
	os.MkdirAll(dir, 0755)
	return filename, DownloadToFile(urlStr, filename)
}

// DownloadToTmpDir 下载指定urlStr的内容到临时目录下，文件名从urlStr末尾获取
func DownloadToTmpDir(urlStr string) (string, error) {
	return DownloadToDir(urlStr, os.TempDir())
}
