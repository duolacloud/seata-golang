package main

import (
	"context"
	"log"
	"os"

	"github.com/duolacloud/seata-golang/pkg/tc/bootstrap"
	"github.com/micro/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "start seata golang tc server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "apollo_namespace",
						Usage: "apollo_namespace",
					},
					&cli.StringFlag{
						Name:  "apollo_address",
						Usage: "apollo_address",
					},
					&cli.StringFlag{
						Name:  "apollo_app_id",
						Usage: "apollo_app_id",
					},
					&cli.StringFlag{
						Name:  "apollo_cluster",
						Usage: "apollo_cluster",
					},
					&cli.StringFlag{
						Name:    "prometheus_addr",
						Usage:   "prometheus_addr",
						EnvVars: []string{"PROMETHEUS_ADDR"},
						Value:   ":16627",
					},
				},
				Action: func(c *cli.Context) error {
					app := bootstrap.App(c)
					if err := app.Start(context.Background()); err != nil {
						log.Fatal(err)
						return err
					}
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
