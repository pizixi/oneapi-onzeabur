package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

const oneAPIURL = "https://github.com/songquanpeng/one-api/releases/download/v0.5.2/one-api"

func main() {
	// 下载 one-api 文件
	err := downloadFile("one-api", oneAPIURL)
	if err != nil {
		fmt.Println("Error downloading one-api:", err)
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

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
