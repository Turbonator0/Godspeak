package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

func randomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func marshal(a interface{}) string {
	marshalled, err := json.Marshal(a)
	check(err)
	var returnString string
	check(json.Unmarshal(marshalled, &returnString))
	return returnString
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func speak(count int) []string {
	var ReturnList []string
	content, err := os.ReadFile("words.json")
	check(err)
	var unmarshalled []interface{}
	err = json.Unmarshal(content, &unmarshalled)
	check(err)
	lastIndex := len(unmarshalled)
	for i := 1; i <= count; i++ {
		item := unmarshalled[randomNumber(lastIndex)]
		ReturnList = append(ReturnList, marshal(item))
	}
	return ReturnList
}
