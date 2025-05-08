package hetzner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/omikkel/whitelist-cf-ips/internal/config"
)

type HetznerPayload struct {
	Rules []HetznerRule `json:"rules"`
}
type HetznerRule struct {
	Direction string `json:"direction"`
	SourceIPs []string `json:"source_ips"`
	Port string `json:"port"`
	Protocol string `json:"protocol"`
	Description string `json:"description"`
}

type Hetzner struct {
	Config *config.Config
}

func InitHetzner(config *config.Config) *Hetzner {
	return &Hetzner{
		Config: config,
	}
}

func (h *Hetzner) WhitelistIPs(ip_range []string) {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", h.Config.HETZNER_API_TOKEN),
		"Content-Type":  "application/json",
	}
	payload := HetznerPayload{
		Rules: []HetznerRule{
			{
				Direction: "in",
				SourceIPs: ip_range,
				Port: "80",
				Protocol: "tcp",
				Description: "Whitelist Cloudflare IPs",
			},
			{
				Direction: "in",
				SourceIPs: ip_range,
				Port: "443",
				Protocol: "tcp",
				Description: "Whitelist Cloudflare IPs",
			},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling hetzner payload:", err)
		return
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.hetzner.cloud/v1/firewalls/%s/actions/set_rules", h.Config.HETZNER_FIREWALL_ID), bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating hetzner POST request:", err)
		return
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending hetzner POST request:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Error sending hetzner POST request: status code", resp.StatusCode)
		return
	}
	log.Println("Successfully whitelisted Cloudflare IPs in Hetzner firewall")
}