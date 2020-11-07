package main

import (
	"lgtool/controller"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "lgtool",
		Usage: "make an explosive entrance",
		Commands: []*cli.Command{
			{
				Name:        "coordtransform",
				Aliases:     []string{"trans", "ct"},
				Usage:       "coordtransform between wgs84,gcj02,bj09",
				Description: "coordtransform and keep 9 places after the decimal point",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "coordinate",
						Aliases: []string{"c"},
						Value:   "w2m",
						Usage: `source coordinate to target coordinate.
								w2g(WGS84toGCJ02),g2w(GCJ02toWGS84) 
								w2b(WGS84toBD09),b2w(BD09toWGS84)
								g2b(GCJ02toBD09),b2g(BD09toGCJ02)`,
						Destination: &controller.TargetCoordinate,
						Required:    true,
					},
				},

				Action: controller.Trans,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
