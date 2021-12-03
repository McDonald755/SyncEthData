package main

import (
	"SyncEthData/db"
	"SyncEthData/log"
	"math"
	"sync"
	"time"
)

func main() {
	log.ConfigLocalFilesystemLogger("./errorLog", "log", time.Hour*24*14, time.Hour*24)
	//cmd.Execute()
	gap := math.Ceil(681.0 / 10.0)
	sw := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		sw.Add(1)
		go db.FixToAccount(int(gap)*i, &sw)
	}
	sw.Wait()

}
