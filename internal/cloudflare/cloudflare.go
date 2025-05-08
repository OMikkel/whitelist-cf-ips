package cloudflare

import (
	"io"
	"log"
	"net/http"
	"strings"
)

type CloudflareIps struct {
	Version string `json:"version"`
	CIDRNetworks []string `json:"cidrs"`
}

func GetIps(version string) *CloudflareIps {
	resp, err := http.Get("https://www.cloudflare.com/ips-v"+version)
	if err != nil {
		log.Println("Error fetching Cloudflare IPs:", err)
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("Error fetching Cloudflare IPs: status code", resp.StatusCode)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil
	}
	if len(body) == 0 {
		log.Println("Error: response body is empty")
		return nil
	}

	cidrs := strings.Split(string(body), "\n")
	
	return &CloudflareIps{
		Version: version,
		CIDRNetworks: cidrs,
	}
}