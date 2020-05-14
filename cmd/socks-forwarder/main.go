package main

import (
	"fmt"
	"github.com/alyakimenko/socks-forwarder/internal/tun"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/alyakimenko/socks-forwarder/internal/handler"
	"github.com/eycorsican/go-tun2socks/core"
	log "github.com/sirupsen/logrus"
)

const (
	version = "0.01"
	mtu     = 1500
)

func main() {
	args := parseFlags()

	if *args.Version {
		fmt.Println(version)
		os.Exit(0)
	}

	switch strings.ToLower(*args.LogLevel) {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		panic("unsupport logging level")
	}

	// Open the tun device.
	dnsServers := strings.Split(*args.TunDNS, ",")
	tunDev, err := tun.OpenTunDevice(*args.TunName, *args.TunAddr, *args.TunGw, *args.TunMask, dnsServers)
	if err != nil {
		log.WithField(
			"TUN Name", *args.TunName,
		).Fatalf("failed to open TUN device: %v", err)
	}
	log.WithFields(log.Fields{
		"TUN Name": *args.TunName,
		"TUN Addr": *args.TunAddr,
		"TUN Gateway": *args.TunGw,
		"TUN Mask": *args.TunMask,
	}).Info("successful open TUN device")

	// Setup TCP/IP stack.
	lwipWriter := core.NewLWIPStack().(io.Writer)

	if err := handler.RegisterSocksHandler(args); err != nil {
		log.Fatal(err)
	}

	if args.DNSFallback != nil && *args.DNSFallback {
		log.Info("Enabled DNS fallback over TCP (overrides the UDP proxy handler).")
		handler.RegisterDNSFallbackHandler()
	}

	// Register an output callback to write packets output from lwip stack to tun
	// device, output function should be set before input any packets.
	core.RegisterOutputFn(func(data []byte) (int, error) {
		written, err := tunDev.Write(data)
		log.Infof("Was written %d bytes to TUN stack", written)

		return written, err
	})

	// Copy packets from tun device to lwip stack, it's the main loop.
	go func() {
		written, err := io.CopyBuffer(lwipWriter, tunDev, make([]byte, mtu))
		log.Infof("Was written %d bytes from lwIP device", written)
		if err != nil {
			log.Fatalf("copying data failed: %v", err)
		}
	}()

	log.Infof("Running tun2socks")

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGHUP)
	<-osSignals
}
