package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)


//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/25


// go run  main.go p 100
// Before
// Action
// 100
// After
func main() {
	app := cli.NewApp()
	app.Name = "SimonWang00"
	app.Version = "1.0.0"
	app.Email = "simon_wang00@163.com"
	app.Commands = []cli.Command{
		{
			Name:    "prot",
			Aliases: []string{"p"},
			Usage:   "用户端口",
			Before: func(c *cli.Context) error {
				fmt.Println("Before")
				return nil
			},
			Action: func(c *cli.Context) {
				fmt.Println("Action")
				fmt.Println(c.Args().First())
			},
			After: func(c *cli.Context) error {
				fmt.Println("After")
				return nil
			},
		},
		{
			Name:    "isdebug",
			Aliases: []string{"p"},
			Usage:   "is debug",
			Before: func(c *cli.Context) error {
				fmt.Println("Before")
				return nil
			},
			Action: func(c *cli.Context) {
				fmt.Println("Action")
				fmt.Println(c.Args().First())
			},
			After: func(c *cli.Context) error {
				fmt.Println("After")
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
