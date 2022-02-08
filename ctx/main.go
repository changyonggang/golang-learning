/*
 * Copyright (c) 2020 KingSoft.com, Inc. All Rights Reserved
 *
 * @file main
 * @author changyonggang(changyonggang@kingsoft.com)
 * @date 2022/2/8 2:50 下午
 * @brief
 *
 */

package main

import (
	"context"
	"fmt"
	"sync"
)

func worker(cancelCtx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(fmt.Sprintf("context value: %v", cancelCtx.Value("key1")))
	for {
		select {
		case val := <-ch:
			fmt.Println(fmt.Sprintf("read from ch value: %d", val))
		case <-cancelCtx.Done():
			fmt.Println("worker is cancelled")
			return
		}
	}
}

func main() {
	rootCtx := context.Background()
	childCtx := context.WithValue(rootCtx, "key1", "value1")
	cancelCtx, cancelFunc := context.WithCancel(childCtx)

	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go worker(cancelCtx, ch, wg)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	cancelFunc()
	wg.Wait()
	close(ch)

}


