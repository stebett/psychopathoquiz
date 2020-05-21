package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Diagnosi struct {
	Diagnosi []Disturbo
}

type Disturbo struct {
	Nome                  string
	Criteri               map[string]string
	Sintomi               Sintomo
	Specificatori         map[string]string
	DiagnosiDifferenziale map[string]string
	Prevalenza            string
	Incidenza             string
	DifferenzaGenere      string
	DifferenzaEta         string
	Decorsi               map[string]string
	Prognosi              string
}

type Sintomo struct {
	Emotivi         map[string]string
	Cognitivi       map[string]string
	Neurovegetativi map[string]string
	Comportamentali map[string]string
}

const fileDiagnosi = "/home/ginko/dev/psychopathoquiz/diagnoquiz/diagnosi.json"

var diagnosi Diagnosi

func loadQuestions() {
	jsonFile, err := os.Open(fileDiagnosi)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &diagnosi)
}
