package main

import (
	"flag"
)

var (
	dir = flag.String("dir", ".", "Directory to serve files from")
)

func main() {
	flag.Parse()
	pcaps, err := glob(*dir, []string{".pcap", ".pcapng"})
	if err != nil {
		panic(err)
	}
	for _, pcap := range pcaps {
		processPcapFile(pcap)
	}
}
