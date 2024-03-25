package main

import (
	"context"
	"fmt"
	"time"
)

func fooFunc(ctx context.Context) {
	fmt.Println("goroutine: 任务开始")
	// 等待Done通道关闭或模拟耗时操作完成
	// select case用法：case中的
	select {
	case <-time.After(5 * time.Second): // 模拟耗时操作
		fmt.Println("goroutine: 任务完成")
	case <-ctx.Done(): // 等待取消信号
		fmt.Println("goroutine: 任务取消")
		return
	}
}

func main() {
	// 创建一个可取消的context
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	// 在goroutine中运行耗时任务
	go fooFunc(ctx)

	// 主程序继续执行...
	time.Sleep(3 * time.Second) // 模拟主程序其他工作耗时
	// fmt.Println("main: 请求取消goroutine任务")
	// cancel() // 发送取消信号

	// 等待goroutine优雅退出（在实际场景中，可能需要同步机制如WaitGroup等待goroutine结束）
	// time.Sleep(1 * time.Second)
	fmt.Println("main: 退出")
}
