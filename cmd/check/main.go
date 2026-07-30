package main

import (
	"fmt"
	"os"
	"log"

	"github.com/urfave/cli/v2"
	"github.com/MehulxBuilds/go-services/internal/checker"
)

func main() {
	app := &cli.App{
		Name: "Website Health Checker",
		Usage: "A tool to check websites health",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "domain",
				Aliases: []string{"d"},
				Usage: "Domain name to check",
				Required: true,
			},
			&cli.StringFlag{
				Name: "port",
				Aliases: []string{"p"},
				Usage: "Port number to check",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if port == "" {
				port = "80"
			}

			domain := c.String("domain")
			if domain == "" {
				log.Fatal("Domain not provided")
			}

			status := checker.Check(domain, port)

			fmt.Println(status)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}