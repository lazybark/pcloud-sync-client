package main

import (
	"github.com/lazybark/pcloud-sync-server/cloud/client"
)

func main() {
	sync := client.NewSyncClient()

	sync.Start()

}
