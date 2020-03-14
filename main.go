package main

import (
	"log"

	"github.com/deeper-x/gocker/gclient"
)

func main() {
	mgr := gclient.NewManager()
	mgr.BuildClient()

	ok := gclient.DownloadImage(mgr, "docker.io/library/alpine")

	if !ok {
		panic("Download error...")
	}

	err := gclient.CreateContainer(mgr)

	if err != nil {
		log.Println(err)
	}
}
