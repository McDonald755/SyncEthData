package main

import (
	"SyncEthData/cmd"
	"SyncEthData/log"
	"time"
)

func main() {
	log.ConfigLocalFilesystemLogger("./log", "log", time.Hour*24*14, time.Hour*24)
	cmd.Execute()
}
