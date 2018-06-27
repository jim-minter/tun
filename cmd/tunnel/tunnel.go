package main

import (
	"github.com/jim-minter/tun/pkg/tunnel"
)

func main() {
	if err := tunnel.Run(); err != nil {
		panic(err)
	}
}
