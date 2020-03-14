package main

import "github.com/deeper-x/gocker/gclient"

func main() {
	mgr := gclient.NewManager()

	mgr.BuildClient()

	mgr.ShowContainers()
}
