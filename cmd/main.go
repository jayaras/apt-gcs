package main

import (
	. "github.com/jayaras/apt-gcs"
)

func main() {
	InitConfig()
	a := AptMethod{}
	a.SendCapabilities()
	a.Run()
}
