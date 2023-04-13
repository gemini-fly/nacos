package dao

import (
	"encoding/json"
	"log"

	"github.com/gemini-fly/nacos/utils"
)

func Config(s string) {
	app := utils.InitNacos

	if err := json.NewDecoder(app).Decode(&s); err != nil {
		log.Fatalln(err)
	}
}
