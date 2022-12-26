package main

import (
	"log"
	"net/http"
	"runtime"
	"time"

	_ "net/http/pprof"
)

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
		time.Sleep(time.Second * 10)
	}()

	runtime.SetBlockProfileRate(1)

}
