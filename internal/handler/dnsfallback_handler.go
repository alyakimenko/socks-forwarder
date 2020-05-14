package handler

import (
	"github.com/eycorsican/go-tun2socks/core"
	"github.com/eycorsican/go-tun2socks/proxy/dnsfallback"
)

func RegisterDNSFallbackHandler() {
	core.RegisterUDPConnHandler(dnsfallback.NewUDPHandler())
}
