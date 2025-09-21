package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	root := context.Background()
	cancelCtx, cancel := context.WithCancel(root)
	defer cancel()

	go func() {
		fmt.Println("Press ENTER to STOP...")
		fmt.Scanln()
		cancel()
	}()

	generateNumbers(cancelCtx)
	fmt.Println("Done")
}

func generateNumbers(ctx context.Context) {
	i := 1
LOOP: // label defined
	for {
		select {
		case <-ctx.Done(): // message will be received here in case when someone presses ENTER
			// stop the processing
			fmt.Println("STOP signal received, shutting down")
			break LOOP
		default:
			fmt.Println(i)
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}
}
