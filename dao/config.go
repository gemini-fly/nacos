package dao

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/gemini-fly/nacos/utils"
)

func Config(s string) {
	app := utils.InitNacos(NamespaceId, DataId, Group)
	if err := json.NewDecoder(bytes.NewBufferString(app)).Decode(&s); err != nil {
		log.Fatalln(err)
	}
}
