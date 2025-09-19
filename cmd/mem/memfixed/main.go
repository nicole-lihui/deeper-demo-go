package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		timer := time.NewTimer(1 * time.Second)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				var m runtime.MemStats
				runtime.ReadMemStats(&m)
				fmt.Printf("[Fixed] Alloc = %v KB, Goroutines = %d\n", m.Alloc/1024, runtime.NumGoroutine())
				timer.Reset(1 * time.Second)
			}
		}
	}()

	select {}
}
