package trace

import (
	"github.com/petermattis/goid"
	"sync"
	"testing"
)

func TestGoRun(t *testing.T) {
	InitReqId()
	t.Log(goid.Get())
	wg := sync.WaitGroup{}
	wg.Add(1)
	t.Log(GetReqId())
	GoRun(func() {
		t.Log(goid.Get())
		t.Log(GetReqId())
		wg.Done()
	})
	wg.Wait()
}
