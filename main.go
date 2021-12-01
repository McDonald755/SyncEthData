package main

import (
	"SyncEthData/cmd"
	"SyncEthData/log"
	"time"
)

func main() {
	log.ConfigLocalFilesystemLogger("./errorLog", "log", time.Hour*24*14, time.Hour*24)
	cmd.Execute()
}
