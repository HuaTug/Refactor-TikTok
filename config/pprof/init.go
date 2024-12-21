package pprof

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
)

func Load() {
	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		if err := http.ListenAndServe(`:6060`, nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
}
