package main

import (
	"fmt"
	"nacos/utils"
)

func main() {
	app := utils.InitNacos("cb93ebba-4410-4828-a95a-9e5d90927c05", "aks", "ops")
	fmt.Println(app)

}
