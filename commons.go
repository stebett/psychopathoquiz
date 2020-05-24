package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"

const bold = "\033[1m"
const underlined = "\033[4m"

func LoadQuestions(filename string, address interface{}) {

	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &address)
}

func PrintLine() {
	fmt.Printf("\n\n________________________________________________________________________________\n\n")
}
