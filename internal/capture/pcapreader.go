package capture

import (
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
)

// PcapReader reads packets from a .pcap file using pcapgo (pure Go).
type PcapReader struct {
	file   *os.File
	reader *pcapgo.Reader
}

// NewPcapReader opens a pcap file for reading.
func NewPcapReader(path string) (*PcapReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open pcap file %q: %w", path, err)
	}

	r, err := pcapgo.NewReader(f)
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("create pcap reader: %w", err)
	}

	return &PcapReader{file: f, reader: r}, nil
}

// Packets returns a gopacket.PacketSource for the file.
func (pr *PcapReader) Packets() *gopacket.PacketSource {
	return gopacket.NewPacketSource(pr.reader, layers.LayerTypeEthernet)
}

// LinkType returns the link layer type for the pcap file.
func (pr *PcapReader) LinkType() layers.LinkType {
	return pr.reader.LinkType()
}

// Close releases the file handle.
func (pr *PcapReader) Close() {
	if pr.file != nil {
		pr.file.Close()
	}
}
