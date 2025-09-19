package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				var m runtime.MemStats
				runtime.ReadMemStats(&m)
				fmt.Printf("[Leak] Alloc = %v KB, Goroutines = %d\n", m.Alloc/1024, runtime.NumGoroutine())
			}
		}
	}()

	select {}
}
