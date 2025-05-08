package main

import (
	"github.com/omikkel/whitelist-cf-ips/internal/cloudflare"
	"github.com/omikkel/whitelist-cf-ips/internal/config"
	"github.com/omikkel/whitelist-cf-ips/internal/hetzner"
)

func main() {
	config := config.GetConfig()

	cloudflare_ipv4 := cloudflare.GetIps("4")
	cloudflare_ipv6 := cloudflare.GetIps("6")
	
	println("Cloudflare IPv4 CIDR Networks:")
	for _, cidr := range cloudflare_ipv4.CIDRNetworks {
		println(cidr)
	}
	println("Cloudflare IPv6 CIDR Networks:")
	for _, cidr := range cloudflare_ipv6.CIDRNetworks {
		println(cidr)
	}

	hetzner := hetzner.InitHetzner(config)

	combined_ips := append(cloudflare_ipv4.CIDRNetworks, cloudflare_ipv6.CIDRNetworks...)

	hetzner.WhitelistIPs(combined_ips)
}