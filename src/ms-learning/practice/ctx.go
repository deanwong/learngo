package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/*
1. 只能编辑 foo 函数
2. foo 必须要调用 slow 函数
3. foo 函数在 ctx 超时后必须立刻返回
4. 【加分项】如果 slow 结束的比 ctx 快，也立刻返回
*/
func main() {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	foo(ctx)
}

func foo(ctx context.Context) {
	ch := make(chan int, 1)
	go func() {
		slow()
		ch <- 1
	}()
	select {
	case <-ctx.Done():
		fmt.Println("ctx timeout quicker than slow")
		return
	case <-ch:
		fmt.Println("slow done quicker than ctx timeout")
		return
	}
}

func slow() {
	n := rand.Intn(3)
	fmt.Printf("sleep %ds\n", n)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Printf("sleep done with %ds\n", n)
}
