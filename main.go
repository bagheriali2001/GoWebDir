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
		},

		Action: func(c *cli.Context) error {
			port := c.Int("port")
			folderPath := c.Args().First()

			if folderPath == "" {
				return fmt.Errorf("Please provide a path to the directory you want to serve")
			}

			fmt.Println("Running the server on port:", port)
			fmt.Println("Serving files from:", folderPath)

			helpers.DisplayAvailableAddresses(strconv.Itoa(port))

			handlerWrapper := handlers.HandlerWrapper{
				RootPath: folderPath,
			}

			// http.HandleFunc("/", handlers.Handler)
			http.HandleFunc("/", handlerWrapper.Handler)

			log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
