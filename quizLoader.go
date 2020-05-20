package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Quizzes struct {
	Quizzes []Quiz
}

type Quiz struct {
	Domanda   string
	Risposte  map[string]string
	Soluzione string
}

const fileDomande = "/home/ginko/dev/psychopathoquiz/domande.json"

func loadQuestions() {

	jsonFile, err := os.Open(fileDomande)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &quizzes)
}
