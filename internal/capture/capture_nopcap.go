//go:build nopcap
// +build nopcap

package capture

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

const (
	DefaultSnapLen = 65535
)

// LiveCapture manages a live packet capture session (stub for systems without libpcap).
type LiveCapture struct {
	iface string
}

// InterfaceInfo describes a network interface.
type InterfaceInfo struct {
	Name        string
	Description string
	Addresses   []string
}

// ListInterfaces returns all available capture interfaces (stub).
func ListInterfaces() ([]InterfaceInfo, error) {
	return []InterfaceInfo{}, nil
}

// NewLiveCapture opens a live capture on the given interface (stub).
func NewLiveCapture(iface, bpfFilter string, snapLen int) (*LiveCapture, error) {
	return nil, fmt.Errorf("live capture is not supported in this environment")
}

// Packets returns a gopacket.PacketSource to iterate packets (stub).
func (lc *LiveCapture) Packets() *gopacket.PacketSource {
	return nil
}

// Interface returns the interface name (stub).
func (lc *LiveCapture) Interface() string {
	return lc.iface
}

// LinkType returns the link layer type for the capture (stub).
func (lc *LiveCapture) LinkType() layers.LinkType {
	return layers.LinkType(layers.LayerTypeEthernet)
}

// Stats returns capture statistics (stub).
func (lc *LiveCapture) Stats() (received, dropped int, err error) {
	return 0, 0, nil
}

// Close stops the capture (stub).
func (lc *LiveCapture) Close() {
}
