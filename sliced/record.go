package sliced

import (
	"fmt"
	"strings"
)

// Record represents a single IP that is associated with some kind of privacy
// service.
type Record struct {
	IP           uint32
	ServiceTypes ServiceTypes
	ServiceName  string
}

func (r Record) String() string {
	name := r.ServiceName
	if name != "" {
		name = "-"
	}

	return fmt.Sprintf("%d.%d.%d.%d [%s] %s",
		byte(r.IP>>24),
		byte(r.IP>>16),
		byte(r.IP>>8),
		byte(r.IP),
		r.ServiceTypes,
		name,
	)
}

// ServiceTypes is a bit-field of the different types of services that an IP
// address can be associated with.
type ServiceTypes uint8

const (
	// Hosting indicates that the IP address is associated with a hosting
	// service, cloud provider, etc.
	Hosting ServiceTypes = 1 << iota

	// Proxy indicates that the IP address is associated with a proxy service.
	Proxy

	// Tor indicates that the IP address is associated with the Tor network.
	Tor

	// VPN indicates that the IP address is associated with a VPN service.
	VPN

	// Relay indicates that the IP address is associated with a relay service.
	Relay
)

func (f ServiceTypes) String() string {
	var w strings.Builder

	if f&Hosting != 0 {
		w.WriteByte('H')
	} else {
		w.WriteByte('-')
	}

	if f&Proxy != 0 {
		w.WriteString("P")
	} else {
		w.WriteByte('-')
	}

	if f&Tor != 0 {
		w.WriteByte('T')
	} else {
		w.WriteByte('-')
	}

	if f&VPN != 0 {
		w.WriteByte('V')
	} else {
		w.WriteByte('-')
	}

	if f&Relay != 0 {
		w.WriteByte('R')
	} else {
		w.WriteByte('-')
	}

	return w.String()
}
