package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "os/exec"
)

const oneAPIURL = "https://github.com/songquanpeng/one-api/releases/download/v0.5.2/one-api"
const oneAPIFileName = "one-api"

func main() {
    if err := runOneAPI(); err != nil {
        fmt.Println("Error:", err)
    }
}

func runOneAPI() error {
    if err := downloadOneAPI(oneAPIFileName, oneAPIURL); err != nil {
        return fmt.Errorf("download error: %w", err)
    }

    if err := os.Chmod(oneAPIFileName, 0755); err != nil {
        return fmt.Errorf("permission error: %w", err)
    }

    cmd := exec.Command("./" + oneAPIFileName)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("execution error: %w", err)
    }

    return nil
}

func downloadOneAPI(filepath string, url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("HTTP GET error: %w", err)
    }
    defer resp.Body.Close()

    out, err := os.Create(filepath)
    if err != nil {
        return fmt.Errorf("file creation error: %w", err)
    }
    defer out.Close()

    if _, err := io.Copy(out, resp.Body); err != nil {
        return fmt.Errorf("file copy error: %w", err)
    }

    return nil
}
