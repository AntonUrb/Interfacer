package clientmodels

import (
	"errors"
	"net"
	"os/exec"
	"strconv"
	"strings"
)

// NetworkInterface represents details about a network interface.
type NetworkInterface struct {
	Name              string   `json:"name"`               // Name of the network interface.
	IPAddresses       []string `json:"ip_addresses"`       // List of IP addresses associated with the interface.
	MACAddress        string   `json:"mac_address"`        // MAC address of the interface.
	MTU               int      `json:"mtu"`                // Maximum Transmission Unit (MTU) of the interface.
	Speed             string   `json:"speed"`              // Speed of the interface.
	Duplex            string   `json:"duplex"`             // Duplex mode of the interface.
	AdminStatus       string   `json:"admin_status"`       // Administrative status of the interface.
	OperationalStatus string   `json:"operational_status"` // Operational status of the interface.
}

// NetworkInterfaces represents a collection of network interfaces.
type NetworkInterfaces struct {
	Interfaces []NetworkInterface `json:"network_interface"` // List of network interfaces.
	Error      string             `json:"error"`             // Error message from the server.
}

// Error represents an error response.
type Error struct {
	Error string `json:"error"` // Error message.
}

// GetInterfaces returns details about all available network interfaces.
func GetInterfaces() ([]NetworkInterface, error) {
	// Get all network interfaces
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var interfaces []NetworkInterface

	// Get IPv4 addresses
	for _, iface := range ifaces {
		addrs, _ := iface.Addrs()
		var ipAddrs []string
		for _, addr := range addrs {
			ipAddrs = append(ipAddrs, strings.Split(addr.String(), "/")[0])
		}

		// Get MAC address
		mac := iface.HardwareAddr.String()

		// Get MTU
		mtu, err := getMTU(iface.Name)
		if err != nil {
			return nil, err
		}

		// Get Speed
		speed, err := getSpeed(iface.Name)
		if err != nil {
			speed = "N/A"
		}

		// Get Duplex
		duplex, err := getDuplex(iface.Name)
		if err != nil {
			duplex = "N/A"
		}

		// Get Admin Status
		adminStatus, err := getAdminStatus(iface.Name)
		if err != nil {
			// If error, set admin status to unknown
			adminStatus = "unknown"
		}

		// Get Operational Status
		operationalStatus, err := getOperationalStatus(iface.Name)
		if err != nil {
			// If error, set operational status to unknown
			operationalStatus = "unknown"
		}

		// Create a NetworkInterface object and append it to the list
		interfaces = append(interfaces, NetworkInterface{
			Name:              iface.Name,
			IPAddresses:       ipAddrs,
			MACAddress:        mac,
			MTU:               mtu,
			Speed:             speed,
			Duplex:            duplex,
			AdminStatus:       adminStatus,
			OperationalStatus: operationalStatus,
		})
	}

	return interfaces, nil
}

// GetInterfaceByName returns the details of a network interface by its name.
func GetInterfaceByName(name string) (*NetworkInterface, error) {
	// Get all network interfaces
	interfaces, err := GetInterfaces()
	if err != nil {
		return nil, err
	}

	// Search for the interface by name
	for _, iface := range interfaces {
		if iface.Name == name {
			return &iface, nil
		}
	}

	// If interface not found, return an error
	return nil, errors.New("there is no such interface")
}

// getMTU retrieves the MTU for a given network interface name.
func getMTU(interfaceName string) (int, error) {
	// Execute the `ip link show` command to get interface details
	cmd := exec.Command("ip", "link", "show", interfaceName)
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	// Parse the output to extract MTU
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "mtu") {
			fields := strings.Fields(line)
			for i, field := range fields {
				if field == "mtu" && i+1 < len(fields) {
					mtuStr := fields[i+1]
					mtu, err := strconv.Atoi(mtuStr)
					if err != nil {
						return 0, err
					}
					return mtu, nil
				}
			}
		}
	}

	return 0, errors.New("MTU not found")
}

// getSpeed retrieves the network speed for a given network interface name using the `ethtool` command.
func getSpeed(interfaceName string) (string, error) {
	// Execute the `ethtool` command
	cmd := exec.Command("ethtool", interfaceName)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Parse the output to extract the network speed
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Speed:") {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				speed := fields[1]
				return speed, nil
			}
		}
	}

	return "", errors.New("speed not found")
}

// getDuplex retrieves the duplex mode of a network interface.
func getDuplex(ifaceName string) (string, error) {
	// Execute the `ethtool` command to get interface details
	cmd := exec.Command("ethtool", ifaceName)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Parse the output to extract the duplex mode value
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Duplex:") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				return parts[1], nil
			}
		}
	}

	return "", errors.New("duplex mode not found")
}

// getAdminStatus returns the administrative status of the given network interface.
func getAdminStatus(ifaceName string) (string, error) {
	// Execute the `ip link show` command to get interface details
	cmd := exec.Command("ip", "link", "show", ifaceName)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Parse the output to extract the administrative status value
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "state") {
			fields := strings.Fields(line)
			for i, field := range fields {
				if field == "state" && i+1 < len(fields) {
					state := fields[i+1]
					if state == "UP" {
						return "enabled", nil
					} else if state == "DOWN" {
						return "disabled", nil
					} else {
						return "unknown", nil
					}
				}
			}
		}
	}

	return "unknown", nil
}

// getOperationalStatus returns the operational status of the given network interface.
func getOperationalStatus(ifaceName string) (string, error) {
	// Execute the `ip link show` command to get interface details
	cmd := exec.Command("ip", "link", "show", ifaceName)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Parse the output to extract the operational status value
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "state") {
			fields := strings.Fields(line)
			for i, field := range fields {
				if field == "state" && i+1 < len(fields) {
					state := fields[i+1]
					if state == "UP" || state == "LOWER_UP" {
						return "UP", nil
					} else if state == "DOWN" {
						return "DOWN", nil
					} else {
						return "unknown", nil
					}
				}
			}
		}
	}
	// If operational status not found, return an error
	return "unknown", nil
}
