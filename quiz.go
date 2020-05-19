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

const domande = "domande.json"
const punteggi = "punteggio.json"
const numeroQuesiti = 166
const numeroEpisodi = 30

var cdf [166]float32

var punteggio int
var clear map[string]func() //create a map for storing clear funcs

func main() {
	rand.Seed(time.Now().UnixNano())

	jsonData := readJson(domande)
	database := caricaDomande(jsonData)

	CallClear()

	fmt.Printf("[!] Iniziamo!\n\n")
	for i := 1; i <= numeroEpisodi; i++ {

		numeroDomanda := ScegliDomanda()
		risultato, risposta := scriviDomanda(database[numeroDomanda])
		punteggio += risultato
		AggiornaPunteggio(numeroDomanda, risultato, punteggi)
		CallClear()
		scriviEsito(risultato, risposta)
	}
	fmt.Printf("Hai totalizzato %d/30 punti!\n", punteggio)

}
