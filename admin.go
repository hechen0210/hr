//go:build ignore

package main

import (
	"fmt"
	"hr/api/admin"
	"hr/config"
	"os"
)

func main() {
	err := config.LoadConfig("./api/admin/config.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	app,err := admin.NewApp()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	app.Run()
}
