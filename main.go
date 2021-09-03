package main

import (
	"os"

	"github.com/masdikaid-rakamin-homework-movie-service/app"
	"github.com/masdikaid-rakamin-homework-movie-service/cli"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())
}
