//go:build ignore

package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// err := config.LoadConfig("./api/company/config.yaml")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// err = config.InitStorage()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// user := request.AdminForm{
	// 	Account:  "admin",
	// 	Name:     "admin",
	// 	Password: "admin",
	// 	Phone:    "",
	// 	Email:    "admin@admin.com",
	// 	Group:    0,
	// 	Status:   1,
	// }
	// s := service.NewAdminService()
	// err = s.CreateOrUpdate(user)
	// fmt.Println(err)
	uuid := uuid.New()
	fmt.Println(uuid)
}
