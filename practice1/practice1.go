/*
Display series of numbers (1,2,3,4, 5....etc) in an infinite loop. The program should quit if someone hits a specific key (Say ESCAPE key).
無限ループで1,2,3...と順番にカウントする。何かキーが押されたら、プログラムを終了させる。
*/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	waitSignal := make(chan os.Signal, 1)
	signal.Notify(waitSignal, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		i := 1
		for {
			fmt.Printf("%d,", i)
			time.Sleep(1 * time.Second)
			i++
		}
	}()

	code := <-waitSignal
	fmt.Print(code)
	os.Exit(0)
}
