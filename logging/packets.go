package logging

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pfring"
)

// Listen will initialize the interface to be used for packet captures,
// then start the loop to handle packets.
func Listen() {
	if ring, err := pfring.NewRing("wlp61s0", 65536, pfring.FlagPromisc); err != nil {
		panic(err)
	} else if err := ring.SetBPFFilter("tcp and port 22"); err != nil {
		panic(err)
	} else if err := ring.Enable(); err != nil {
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(ring, layers.LinkTypeEthernet)

		for packet := range packetSource.Packets() {
			printPacket(packet)
		}
	}
}

func printPacket(packet gopacket.Packet) {
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		fmt.Printf("Received packet from %d", tcp.SrcPort)
	}
}
