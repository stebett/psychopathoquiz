package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Quiz struct {
	Domanda   string
	Risposte  map[string]string
	Soluzione string
}

const filename = "domande.json"
const numeroQuesiti = 166
const numeroEpisodi = 30

var punteggio int
var clear map[string]func() //create a map for storing clear funcs

func main() {
	rand.Seed(time.Now().UnixNano())

	jsonData := readJson()
	database := caricaDomande(jsonData)

	CallClear()

	fmt.Printf("[!] Iniziamo!\n\n")
	for i := 1; i <= numeroEpisodi; i++ {

		risultato, risposta := scriviDomanda(database[rand.Intn(numeroQuesiti)])
		punteggio += risultato
		CallClear()
		scriviEsito(risultato, risposta)
	}
	fmt.Printf("Hai totalizzato %d/30 punti!\n", punteggio)

}
