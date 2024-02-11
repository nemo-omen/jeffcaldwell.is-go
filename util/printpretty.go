package util

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrintPretty(s interface{}) {
	sJson, err := json.MarshalIndent(s, "", "  ")

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%s\n", sJson)
}
