//go:build ignore

package main

import (
	"fmt"
	"hr/api/company"
	"hr/config"
	"os"
)

func main() {
	err := config.LoadConfig("./api/company/config.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	app,err := company.NewApp()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	app.Run()
}
