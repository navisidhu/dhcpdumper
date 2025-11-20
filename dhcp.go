package main

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func processPcapFile(filename string) error {
	handle, err := pcap.OpenOffline(filename)
	if err != nil {
		return fmt.Errorf("OpenOffline: %w", err)
	}
	defer handle.Close()

	// Optional: restrict to UDP DHCP ports via BPF to reduce workload
	// DHCPv4: 67/68, DHCPv6: 546/547
	if err := handle.SetBPFFilter("udp and (port 67 or port 68 or port 546 or port 547)"); err != nil {
		return fmt.Errorf("SetBPFFilter: %w", err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		printDHCPv4(packet)
	}

	return nil
}

func printDHCPv4(packet gopacket.Packet) {
	// for _, b := range packet.Data() {
	// 	fmt.Printf("0x%02x,", b)
	// }
	// fmt.Println()

	dhcpLayer := packet.Layer(layers.LayerTypeDHCPv4)
	if dhcpLayer == nil {
		return
	}

	fmt.Println("{")
	for _, b := range dhcpLayer.LayerContents() {
		fmt.Printf("0x%02x,", b)
	}
	fmt.Println("},")

}
