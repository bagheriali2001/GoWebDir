package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/bagheriali2001/GoWebDir/handlers"
	"github.com/bagheriali2001/GoWebDir/helpers"
)

// TODO: add json api version

func main() {
	app := &cli.App{
		Name:      "go-web-dir",
		Usage:     "A simple web server to serve files in a directory",
		UsageText: "contrive - demonstrating the available API",
		ArgsUsage: "<path>",

		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Value:       3000,
				Usage:       "Port to run the server on",
				DefaultText: "3000",
				Action: func(c *cli.Context, port int) error {
					if (port < 1024 && (port != 80 || port != 443)) || port > 65535 {
						return fmt.Errorf("Invalid port number: %d. Please provide a valid port (Check --help for more information)", port)
					}

					return nil
				},
			},
			&cli.BoolFlag{
				Name:        "show-hidden-files",
				Value:       false,
				Usage:       "Show hidden files in the directory",
				DefaultText: "false",
			},
			&cli.BoolFlag{
				Name:        "show-hidden-folders",
				Value:       false,
				Usage:       "Show hidden folders in the directory",
				DefaultText: "false",
			},
		},

		Action: func(c *cli.Context) error {
			// TODO: Add a configuration summery log

			port := c.Int("port")
			folderPath := c.Args().First()

			helpers.PrintStartupConfig(c.Int("port"), c.Args().First(), c.Bool("show-hidden-files"), c.Bool("show-hidden-folders"))

			if folderPath == "" {
				return fmt.Errorf("Please provide a path to the directory you want to serve")
			}

			helpers.DisplayAvailableAddresses(strconv.Itoa(port))

			handlerWrapper := handlers.HandlerWrapper{
				RootPath:          folderPath,
				ShowHiddenFiles:   c.Bool("show-hidden-files"),
				ShowHiddenFolders: c.Bool("show-hidden-folders"),
			}

			http.HandleFunc("/", handlerWrapper.Handler)

			log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
