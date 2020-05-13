package config

import "time"

// Config defines command line arguments
type Config struct {
	Version         *bool
	TunName         *string
	TunAddr         *string
	TunGw           *string
	TunMask         *string
	TunDNS          *string
	TunPersist      *bool
	BlockOutsideDNS *bool
	ProxyType       *string
	ProxyServer     *string
	ProxyHost       *string
	ProxyPort       *uint16
	UDPTimeout      *time.Duration
	LogLevel        *string
	DNSFallback     *bool
}
