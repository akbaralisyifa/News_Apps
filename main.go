package main

import (
	"fmt"
	"newsapps/configs"
)

func main() {
	setup := configs.ImportSetting()
	connection, err := configs.ConnectDB(setup)
	if err != nil {
		fmt.Println("Stop program, masalah pada database", err.Error())
		return
	}
	fmt.Println("testing BERHASIL!", connection)
}
