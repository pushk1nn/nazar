package logging

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// Listen will initialize the interface to be used for packet captures,
// then start the loop to handle packets.
func Listen() {
	if handle, err := pcap.OpenLive("wlp61s0", 1600, true, pcap.BlockForever); err != nil {
		panic(err)
	} else if err := handle.SetBPFFilter("tcp and port 80"); err != nil {
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
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		fmt.Printf("Received packet from %d", tcp.SrcPort)
	}
}
