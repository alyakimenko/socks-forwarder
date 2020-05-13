package main

import (
	"flag"

	"github.com/alyakimenko/socks-forwarder/config"
)

func parseFlags() *config.Config {
	c := &config.Config{}

	c.Version = flag.Bool("version", false, "Print version")
	c.TunName = flag.String("tunName", "tun1", "TUN interface name")
	c.TunAddr = flag.String("tunAddr", "10.255.0.2", "TUN interface address")
	c.TunGw = flag.String("tunGw", "10.255.0.1", "TUN interface gateway")
	c.TunMask = flag.String("tunMask", "255.255.255.0", "TUN interface netmask, it should be a prefixlen (a number) for IPv6 address")
	c.TunDNS = flag.String("tunDns", "8.8.8.8,8.8.4.4", "DNS resolvers for TUN interface (only need on Windows)")
	c.TunPersist = flag.Bool("tunPersist", false, "Persist TUN interface after the program exits or the last open file descriptor is closed (Linux only)")
	c.BlockOutsideDNS = flag.Bool("blockOutsideDns", false, "Prevent DNS leaks by blocking plaintext DNS queries going out through non-TUN interface (may require admin privileges) (Windows only) ")
	c.ProxyType = flag.String("proxyType", "socks", "Proxy handler type")
	c.LogLevel = flag.String("loglevel", "info", "Logging level. (debug, info, warn, error)")

	flag.Parse()
	return c
}
