package dao

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/gemini-fly/nacos/utils"
)

const (
	NamespaceId = "dcc3ab24-6834-4cd1-bd57-2c24d61ffe5f"
	DataId      = "ak"
	Group       = "ops"
)

func Config(s string) {
	app := utils.InitNacos(NamespaceId, DataId, Group)
	if err := json.NewDecoder(bytes.NewBufferString(app)).Decode(&s); err != nil {
		log.Fatalln(err)
	}
}
