package main

import (
	"SyncEthData/log"
	"SyncEthData/cmd"
	"time"
)

func main() {
	log.ConfigLocalFilesystemLogger("./log", "log", time.Hour*24*14, time.Hour*24)
	cmd.Execute()
	//cmd.ScanBlock(0)
}
