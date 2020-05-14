package handler

import (
	"errors"
	"fmt"
	"net"

	"github.com/alyakimenko/socks-forwarder/config"
	"github.com/eycorsican/go-tun2socks/core"
	"github.com/eycorsican/go-tun2socks/proxy/socks"
	log "github.com/sirupsen/logrus"
)

func RegisterSocksHandler(args *config.CmdArgs) error {
	if args.ProxyServer == nil {
		return errors.New("proxy server not set")
	}
	// Verify proxy server address.
	proxyAddr, err := net.ResolveTCPAddr("tcp", *args.ProxyServer)
	if err != nil {
		return fmt.Errorf("invalid proxy server address: %v", err)
	}
	proxyHost := proxyAddr.IP.String()
	proxyPort := uint16(proxyAddr.Port)

	log.WithFields(log.Fields{
		"Proxy Addr": *args.ProxyServer,
		"Proxy Host": *args.ProxyHost,
		"Proxy Port": *args.ProxyPort,
	}).Info("Proxy details")

	core.RegisterTCPConnHandler(socks.NewTCPHandler(proxyHost, proxyPort))
	core.RegisterUDPConnHandler(socks.NewUDPHandler(proxyHost, proxyPort, *args.UDPTimeout))

	return nil
}
