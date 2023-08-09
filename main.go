package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
)

//go:embed one-api
var oneAPI []byte

func main() {
	// 解压缩 one-api 到当前目录
	err := os.WriteFile("one-api", oneAPI, 0755)
	if err != nil {
		fmt.Println("Error writing one-api to current directory:", err)
		return
	}
	defer os.Remove("one-api")

	// 赋予权限
	err = os.Chmod("one-api", 0755)
	if err != nil {
		fmt.Println("Error setting permissions for one-api:", err)
		return
	}

	// 运行 one-api
	cmd := exec.Command("./one-api")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running one-api:", err)
		return
	}
}
