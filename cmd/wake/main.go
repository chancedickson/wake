package main

import (
	"fmt"
	"os"

	"github.com/pheebcodes/wake/pkg/config"
	"github.com/pheebcodes/wake/pkg/wol"
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
	err = wol.SendMagicPacket(config.IPAddress, config.Port, config.MacAddress)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
