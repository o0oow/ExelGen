package configs

import (
	"encoding/json"
	"log"
	"os"
	"sportsman/models"
)

var Config models.Config

func Setup(F string) {
	byteValue, err := os.ReadFile(F)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = json.Unmarshal(byteValue, &Config)
	if err != nil {
		log.Fatal(err)
		return
	}
}
