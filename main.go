package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

const oneAPIURL = "https://github.com/songquanpeng/one-api/releases/download/v0.5.2/one-api"

func main() {
	// 创建一个默认的 Gin 路由引擎
	router := gin.Default()

	// 设置转发规则
	targetURL, _ := url.Parse("http://localhost:3000")
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	router.Any("/*path", func(c *gin.Context) {
		go runOneapi()
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	// 启动 HTTP 服务，监听 8080 端口
	router.Run(":8080")

}
func runOneapi() {
	// 下载 one-api 文件
	err := downloadFile("one-api", oneAPIURL)
	if err != nil {
		fmt.Println("Error downloading one-api:", err)
		return
	}

	// 赋予权限
	err = os.Chmod("one-api", 0755)
	if err != nil {
		fmt.Println("Error setting permissions for one-api:", err)
		return
	}

	// 运行 one-api
	cmd := exec.Command("./one-api", "--port", "3000")
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
