package main

import (
	"log"

	"github.com/deeper-x/gocker/gclient"
)

func main() {
	mgr := gclient.NewManager()

	mgr.BuildClient()

	log.Println(gclient.ShowImagesIDs(mgr))
}
