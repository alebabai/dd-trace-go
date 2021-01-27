//+build ignore

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).

package main

import (
	"os"
	"runtime/pprof"
)

func main() {
	stackDump(100)
}

func stackDump(remaining int) {
	if remaining > 0 {
		stackDump(remaining - 1)
	} else {
		pprof.Lookup("goroutine").WriteTo(os.Stdout, 2)
	}
}
