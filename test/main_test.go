package test

import (
	"github.com/gannicus-w/go_keepAlive/lib"
	"testing"
)

func TestGoKeepAlive(t *testing.T) {
	wm := lib.NewWorkerManager(10)

	wm.StartWorkerPool()
}
