package cmm

import (
	"archive/tar"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/pkg/errors"
)

// Md5Hex 获取对bs进行md5后的hex值
func Md5Hex(bs []byte) string {
	h := md5.New()
	_, err := h.Write(bs)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(h.Sum(nil))
}

// GetWorkDir 获取当前工作目录
func GetWorkDir() string {
	ex, err := os.Executable()
	if err != nil {
		return "."
	}
	return path.Dir(ex)
}

// GetPhysicsInterfaces 获取当前机器的物理网卡
func GetPhysicsInterfaces() []net.Interface {
	// 记录当前的物理网卡
	phs := make([]net.Interface, 0)
	faces, _ := net.Interfaces()
	// 过滤出物理网卡
	for _, i := range faces {
		if i.Flags == 0 || len(i.Name) == 0 || i.MTU == 0 {
			continue
		}

		// 名称过滤，支持em*, eth*
		if !strings.HasPrefix(i.Name, "em") && !strings.HasPrefix(i.Name, "eth") {
			continue
		}

		// 检查硬件地址第二位十六进制数，如果不为偶数则不是单播地址
		if len(i.HardwareAddr.String()) < 2 || !strings.Contains("02468ACEace", string(i.HardwareAddr.String()[1])) {
			continue
		}

		if (i.Flags & net.FlagPointToPoint & net.FlagLoopback) > 0 {
			continue
		}

		phs = append(phs, i)

		// addrs, err := i.Addrs()
		// if err != nil {
		// 	logrus.Errorln("net.interfaces.addrs return Error: " + err.Error())
		// } else {

		// 	fmt.Println(i, addrs)
		// 	for _, a := range addrs {
		// 		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
		// 			ip4 := ipnet.IP.To4()
		// 			if ip4 == nil {
		// 				continue
		// 			}

		// 			phs = append(phs, i)
		// 		}
		// 	}
		// }

	}

	return phs
}

// NohupRun 以nohup 形式执行命令
func NohupRun(cmd string) error {
	nohup := exec.Command("sh", "-c", fmt.Sprintf("nohup %s &", cmd))
	if nohup == nil {
		return errors.New("make nohup command failed")
	}

	err := nohup.Run()
	if err != nil {
		return errors.Wrap(err, "run hohup failed")
	}
	return nil
}

// UnZipTar 解压tar.gz文件。成功返回解压出的顶级文件夹名称
func UnZipTar(filename, dstFolder string) (string, error) {
	// file read
	fr, err := os.Open(filename)
	if err != nil {
		return "", errors.Wrap(err, "open file failed")
	}
	defer fr.Close()

	// gzip read
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return "", errors.Wrap(err, "read file failed")
	}
	defer gr.Close()

	err = os.MkdirAll(dstFolder, 0755)
	if err != nil {
		return "", errors.Wrap(err, "make directory for destination folder failed")
	}

	// tar read
	tr := tar.NewReader(gr)

	// 读取文件
	rootDirName := ""
	for {
		var h *tar.Header
		h, err = tr.Next()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			err = errors.Wrap(err, "read from zip file failed")
			break
		}

		name := path.Clean(h.Name)

		if h.FileInfo().IsDir() {
			if len(rootDirName) == 0 {
				rootDirName = name
			}
			err = os.MkdirAll(dstFolder+"/"+name+"/", 0755)
			if err != nil {
				err = errors.Wrap(err, "make directory for unzip file failed")
				break
			}
			continue
		}

		// 写文件
		var fw *os.File
		fw, err = os.Create(dstFolder + "/" + name)
		if err != nil {
			err = errors.Wrap(err, "create unzip file failed")
			break
		}
		_, err = io.Copy(fw, tr)
		fw.Close()
		if err != nil {
			err = errors.Wrap(err, "write unzip file failed")
			break
		}
	}

	if err != nil {
		return "", err
	}
	return rootDirName, err
}
