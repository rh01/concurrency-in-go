package goroutines

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}

	receive := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receive()
	// begin the performance timer
	b.StartTimer()
	// tell the two gorountines to begin.
	close(begin)
	wg.Wait()
}

// RUN
// go test -bench=. -cpu=1 src/mycode/0922-currency/goroutine-benchmark/fig-ctx-switch_test.go
