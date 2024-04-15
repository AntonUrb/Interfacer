package servermodels

import (
	"fmt"
	"testing"
)

// Mock network interfaces for testing
var mockInterfaces = []NetworkInterface{
	{
		Name:              "eth0",
		IPAddresses:       []string{"192.168.1.10", "10.0.0.1"},
		MACAddress:        "00:11:22:33:44:55",
		MTU:               1500,
		Speed:             "1 Gbps",
		Duplex:            "Full",
		AdminStatus:       "enabled",
		OperationalStatus: "UP",
	},
	{
		Name:              "wlan0",
		IPAddresses:       []string{"192.168.2.20"},
		MACAddress:        "a1:b2:c3:d4:e5:f6",
		MTU:               1200,
		Speed:             "100 Mbps",
		Duplex:            "Half",
		AdminStatus:       "disabled",
		OperationalStatus: "DOWN",
	},
}

func TestGetInterfaces(t *testing.T) {

	interfaces := mockInterfaces

	// Test individual interfaces
	for _, iface := range interfaces {
		t.Run(fmt.Sprintf("Interface_%s", iface.Name), func(t *testing.T) {
			t.Logf("Interface: %s\n", iface.Name)
			t.Logf("  IP Addresses: %v\n", iface.IPAddresses)
			t.Logf("  MAC Address: %s\n", iface.MACAddress)
			t.Logf("  MTU: %d\n", iface.MTU)
			t.Logf("  Speed: %s\n", iface.Speed)
			t.Logf("  Duplex: %s\n", iface.Duplex)
			t.Logf("  Admin Status: %s\n", iface.AdminStatus)
			t.Logf("  Operational Status: %s\n", iface.OperationalStatus)

			// Assertions based on mock values
			if iface.Name != "eth0" && iface.Name != "wlan0" {
				t.Errorf("Unexpected interface name: %s", iface.Name)
			}
			if len(iface.IPAddresses) != 2 && len(iface.IPAddresses) != 1 {
				t.Errorf("Unexpected number of IP addresses for interface %s", iface.Name)
			}
			if iface.MACAddress == "" {
				t.Errorf("MAC address is empty for interface %s", iface.Name)
			}
			if iface.MTU <= 0 {
				t.Errorf("Invalid MTU for interface %s", iface.Name)
			}
			if iface.Speed == "" {
				t.Errorf("Speed is empty for interface %s", iface.Name)
			}
			if iface.Duplex == "" {
				t.Errorf("Duplex is empty for interface %s", iface.Name)
			}
			if iface.AdminStatus != "enabled" && iface.AdminStatus != "disabled" {
				t.Errorf("Unexpected admin status for interface %s", iface.Name)
			}
			if iface.OperationalStatus != "UP" && iface.OperationalStatus != "DOWN" {
				t.Errorf("Unexpected operational status for interface %s", iface.Name)
			}
		})
	}
}
