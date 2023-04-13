package dao

import (
	"encoding/json"
	"log"
	"nacos/utils"
)

func Config(s string) {
	app := utils.InitNacos(NamespaceId, DataId, Group)
	if err := json.NewDecoder(app).Decode(&s); err != nil {
		log.Fatalln(err)
	}
}
