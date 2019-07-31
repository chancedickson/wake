package main

import (
	"fmt"
	"os"

	"github.com/chancedickson/wake/pkg/config"
	"github.com/chancedickson/wake/pkg/wol"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [name of computer to start]\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	config, err := config.LoadConfig(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	wol.SendMagicPacket(config.IPAddress, config.Port, config.MacAddress)
}
