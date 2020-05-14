package config

import (
	"time"
)

// Config defines command line arguments
type CmdArgs struct {
	Version *bool
	// TunName is TUN interface name
	TunName *string
	// TunAddr is TUN interface address
	TunAddr *string
	// FunGw is TUN interface gateway
	TunGw *string
	// TunMask is TUN interface netmask, it should be a prefixlen (a number) for IPv6 address
	TunMask *string
	// TunDNS is DNS resolvers for TUN interface (only need on Windows)
	TunDNS *string
	// TunPersist persists TUN interface after the program exits
	// or the last open file descriptor is closed (Linux only)
	TunPersist *bool
	// BlockOutsideDNS prevents DNS leaks by blocking plaintext DNS queries going out
	// through non-TUN interface (may require admin privileges) (Windows only)
	BlockOutsideDNS *bool
	// ProxyType is proxy handler type
	ProxyType   *string
	// ProxyServer is proxy server address
	ProxyServer *string
	// ProxyHost is proxy host
	ProxyHost   *string
	// ProxyHost is proxy port
	ProxyPort   *uint16
	UDPTimeout  *time.Duration
	// LogLevel is logging level. (debug, info, warn, error)
	LogLevel    *string
	DNSFallback *bool
}
