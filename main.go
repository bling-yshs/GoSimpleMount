package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        filePath := "." + r.URL.Path
        http.ServeFile(w, r, filePath)
    })

    server := &http.Server{Addr: ":6421"}

    go func() {
        fmt.Println("文件服务器运行在 6421 端口")
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            fmt.Printf("listen: %+s\n", err)
        }
    }()

    // 等待stop命令
    var cmd string
    for {
        fmt.Scanln(&cmd)
        if cmd == "stop" {
            break
        }
    }

    // 关闭服务器
    if err := server.Close(); err != nil {
        fmt.Printf("Server Close: %+s\n", err)
    }
}
