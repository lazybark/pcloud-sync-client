package main

import (
	"github.com/lazybark/pcloud-sync-server/sync"
)

func main() {
	sync := sync.NewSyncClient()

	sync.Start()

}
