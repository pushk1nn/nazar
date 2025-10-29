package logging

import (
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// device gets the environment variable to listen on
var device = os.Getenv("IFACE")

// Listen will initialize the interface to be used for packet captures,
// then start the loop to handle packets.
func Listen() {
	if handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever); err != nil {
		panic(err)
	} else if err := handle.SetBPFFilter("tcp src port 22"); err != nil { // For early development, listen for ssh
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

		for packet := range packetSource.Packets() {
			printPacket(packet)
		}
	}
}

// printPacket prints the packet's source port.
func printPacket(packet gopacket.Packet) {
	// Print source IP of packet
	if ipv4Layer := packet.Layer(layers.LayerTypeIPv4); ipv4Layer != nil {
		ipv4, _ := ipv4Layer.(*layers.IPv4)
		fmt.Println("Source IP:", ipv4.SrcIP)
	}
}
