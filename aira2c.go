package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	var cmd *exec.Cmd
	var err error

	cmd = exec.Command("aria2c", "--enable-rpc", "--rpc-listen-all")
	if err = cmd.Start(); err != nil {
		fmt.Printf("打开 aria2 失败，错误信息：%v\n请将 aria2 加入到 PATH 环境变量中。\naria2 下载地址 https://aria2.github.io \n", err)
		os.Exit(1)
	}

	fmt.Println("aria2 已经启动，成功开启 RPC 服务...")

	if !openBrowser("index.html") {
		fmt.Printf("没有找到 webui-aria2，错误信息：%v\n请将本程序放到 webui-aria2 目录下。\nwebui-aria2 下载地址 https://github.com/ziahamza/webui-aria2 \n", err)
		os.Exit(1)
	}
}

// openBrowser tries to open the URL in a browser,
// and returns whether it succeed in doing so.
func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}
