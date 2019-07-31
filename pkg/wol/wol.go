package wol

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net"
	"strings"
)

var packetHeader = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

func removeCharacters(s string, cs []string) string {
	for _, c := range cs {
		s = strings.ReplaceAll(s, c, "")
	}
	return s
}

func macAddressToBytes(macAddress string) ([]byte, error) {
	joinedMacAddress := removeCharacters(macAddress, []string{"-", ":"})
	slice, err := hex.DecodeString(joinedMacAddress)
	if err != nil {
		return nil, err
	}
	return slice, nil
}

// SendMagicPacket sends a magic packet of the given hardware address to the given IP address at the given port
func SendMagicPacket(ip string, port uint16, macAddress string) error {
	macAddressBytes, err := macAddressToBytes(macAddress)
	if err != nil {
		return err
	}
	repeatedMac := bytes.Repeat(macAddressBytes, 16)
	packet := bytes.Join([][]byte{packetHeader, repeatedMac}, []byte{})
	conn, err := net.Dial("udp", fmt.Sprintf("%v:%v", ip, port))
	defer conn.Close()
	if err != nil {
		return err
	}
	_, err = conn.Write(packet)
	if err != nil {
		return err
	}
	return nil
}
