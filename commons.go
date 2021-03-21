package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"

const bold = "\033[1m"
const underlined = "\033[4m"

func ChooseGame() {
	PrintLine()
	fmt.Println("Choose your game:\n[1] GuessCross\n[2] GuessName\n[3] GuessSpec\n[4] GuessSint\n[5] Exit")
	answer := Input()

	if answer[0] == "1" {
		GuessCross()
	}
	if answer[0] == "2" {
		GuessName()
	}
	if answer[0] == "3" {
		GuessSpec()
	}
	if answer[0] == "4" {
		GuessSint()
	}
	if answer[0] == "5" {
		os.Exit(0)
	}
}

func LoadJson(filename string, address interface{}) {

	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &address)
}

func Input() []string {
	fmt.Print("\n[>] ")
	var answer string
	fmt.Scanln(&answer)

	answerArray := FormatString(answer)
	return answerArray
}

func FormatString(s string) []string {
	s = strings.ToLower(s)
	s = strings.TrimSuffix(s, "\n")
	sArray := strings.Split(s, " ")

	return sArray
}
func PrintLine() {
	fmt.Printf("\n\n________________________________________________________________________________\n\n")
}
