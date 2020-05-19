package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

type Punteggi struct {
	Punteggi []Punteggio
}

type Punteggio struct {
	Corrette  float32
	Sbagliate float32
}

const filePunteggi = "punteggio.json"

func ScegliDomanda() int {
	ComputeCDF()
	fmt.Println(cdf)
	r := rand.Float32()

	domanda := 0
	for r > cdf[domanda] {
		domanda++
	}
	return domanda
}

func AggiornaPunteggio(numeroDomanda int, risultato int, punteggi Punteggi) {
	return

}

func ComputeCDF() {
	pointer := &cdf
	loadScore()
	var normalizer float32

	for i := range cdf {
		pointer[i] = punteggi.Punteggi[i].Sbagliate / punteggi.Punteggi[i].Corrette
		normalizer += pointer[i]
	}

	for i := range cdf {
		pointer[i] = pointer[i] / normalizer
	}

	for i := range cdf {
		if i > 0 {
			pointer[i] += pointer[i-1]
		}
	}
}

func loadScore() {

	jsonFile, err := os.Open(filePunteggi)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &punteggi)
}
