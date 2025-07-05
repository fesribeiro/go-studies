package app

import (
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate creates a new CLI application for finding IP addresses.
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "IP Finder"
	app.Usage = "A simple CLI tool to find your IP address"
	app.Version = "1.0.0"

	flags := []cli.Flag{
			cli.StringFlag{
				Name:  "host",
				Value: "google.com.br",
			},
		}

	app.Commands = []cli.Command{
		{
			Name:    "find-ip",
			Usage:   "Find your public IP address",
			Flags: flags,
			// Aliases: []string{"f"},
			Action: findIp,
		},
		{
			Name: "servers",
			Usage: "Find DNS servers for a domain",
			Flags: flags,
			Action: findServers,
		},
	}

	return app
}

func findServers(c *cli.Context) {
	host := c.String("host")
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal("Error looking up DNS servers:", err)
	}
	// Log the found DNS servers
	for _, server := range servers {
		log.Println("Found DNS server:", server.Host)
	}
}

func findIp(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal("Error looking up IP address:", err)
	}
	// Log the found IP addresses
	for _, ip := range ips {
		log.Println("Found IP address:", ip)
	}
}	