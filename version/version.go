// Package version 用于 go build 时，通过-ldflags “-X”标识来指定当期那构建的版本号
// 示例makefile如下：
/*
#!/bin/bash
version=$1
function build()
{
    # 准备好要填入程序内的信息
    buildtime=`date +%Y-%m%d_%I:%M:%S`
    githash=`git rev-parse HEAD`
    if [ "$versionno" == "" ]; then
        version=`git tag --contains="$githash"`
	fi
	# version包路径
	vpkg="github.com/monotone/tools/version"

    GOOS=linux GOARCH=amd64 go build -v \
        -ldflags "-s -w -X $vpkg.BuildTime=$buildtime -X $vpkg.GitHash=$githash -X $vpkg.VersionNo=$version" \
        -o ./my-program ./
}

build || { echo "构建失败！"; exit 1; }
*/

package version

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

var (
	// BuildTime 指定构建时间
	BuildTime = ""
	// GitHash 对应构建时的tag id 或者 commit id
	GitHash = ""
	// Version 版本号
	Version = ""
)

// PrintVersionInfo 打印出版本相关信息到ws，如果没有提供ws,则输出到stdout
func PrintVersionInfo(ws ...io.Writer) {
	writer := io.Writer(os.Stdout)
	if len(ws) > 0 {
		writer = ws[0]
	}
	writer.Write([]byte(fmt.Sprintln("Build time\t\t:", BuildTime)))
	writer.Write([]byte(fmt.Sprintln("Hash code \t\t:", GitHash)))
	writer.Write([]byte(fmt.Sprintln("Version number\t\t:", Version)))
	writer.Write([]byte(fmt.Sprintln("Go Version Info\t\t:", runtime.Version())))
}
