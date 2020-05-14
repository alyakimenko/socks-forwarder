package handler

import (
	"fmt"
	"net"

	"github.com/alyakimenko/socks-forwarder/config"
	"github.com/eycorsican/go-tun2socks/core"
	"github.com/eycorsican/go-tun2socks/proxy/socks"
)

func RegisterSocksHandler(args *config.CmdArgs) error {
	// Verify proxy server address.
	proxyAddr, err := net.ResolveTCPAddr("tcp", *args.ProxyServer)
	if err != nil {
		return fmt.Errorf("invalid proxy server address: %v", err)
	}
	proxyHost := proxyAddr.IP.String()
	proxyPort := uint16(proxyAddr.Port)

	core.RegisterTCPConnHandler(socks.NewTCPHandler(proxyHost, proxyPort))
	core.RegisterUDPConnHandler(socks.NewUDPHandler(proxyHost, proxyPort, *args.UDPTimeout))

	return nil
}
