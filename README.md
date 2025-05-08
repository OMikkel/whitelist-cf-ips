# Update Hetzner Firewall Rules with Cloudflare CIDR Ranges

A simple script to get the latest cloudflare CIDR ranges and update the hetzner firewall rules

# Setup cron job

1. Run `crontab -e` to edit the cron jobs.
2. Add the following line to run the script every day at 0 AM:
   ```
   0 0 * * * <your shell> <path to your binary executable>
   ```

## Supported architectures

Currently it is built for the following architectures:

- `linux/amd64`
- `macos/arm64`
