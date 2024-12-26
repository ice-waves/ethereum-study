package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done(): // 检测到取消信号
			fmt.Printf("Worker %d: stopping\n", id)
			return
		default:
			fmt.Printf("Worker %d: working\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println(time.Now().Format(time.DateTime))

	// 启动两个 goroutine
	go worker(ctx, 1)
	go worker(ctx, 2)
	// 让主线程运行 3 秒钟
	time.Sleep(3 * time.Second)
	fmt.Println(time.Now().Format(time.DateTime))

	// 取消 context，通知所有 goroutine 停止
	fmt.Println("Canceling context...")
	cancel()

	// 等待 goroutine 停止
	time.Sleep(2 * time.Second)
}
