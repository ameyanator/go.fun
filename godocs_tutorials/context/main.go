package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	deadLine := time.Now().Add(1500 * time.Millisecond)
	ctx, cancelCtx := context.WithDeadline(ctx, deadLine)
	defer cancelCtx()

	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		select {
		case printCh <-num:
			time.Sleep(1*time.Second)
		case <-ctx.Done():
			break;
		}
	}
	cancelCtx()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("doSomething finished")
}

func doAnother(ctx context.Context, printChan <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Println("doAnother err", err.Error())
			}
			fmt.Println("doAnother finished")
			return
		case num := <-printChan:
			fmt.Println("doAnother", num)
		}

	}
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "nothing")
	doSomething(ctx)
}
