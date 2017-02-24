package main

import "fmt"
import "time"

func main() {
	step := 10*time.Microsecond
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println(-i)
			time.Sleep(step)
		}
	}()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(step)
	}
	time.Sleep(step)
	fmt.Println("done")
}
