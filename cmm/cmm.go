package cmm

import (
	"crypto/md5"
	"encoding/hex"
	"net"
	"os"
	"path"
	"strings"
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
