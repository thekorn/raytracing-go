//go:build pprof
// +build pprof

package main

import (
	l "log"
	"os"
	"runtime/pprof"
)

func init() {
	f, perr := os.Create("./tmp/cpu.pprof")
	if perr != nil {
		l.Fatal(perr)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
}
