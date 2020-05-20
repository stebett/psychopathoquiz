package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numeroQuesiti = 166
const numeroEpisodi = 30

var cdf [166]float32
var quizzes Quizzes
var punteggi Punteggi

func main() {
	rand.Seed(time.Now().UnixNano())

	loadQuestions()

	var punteggio int
	var numeroDomanda int
	for i := 1; i <= numeroEpisodi; i++ {

		numeroDomanda = ScegliDomanda()
		risultato, risposta := scriviDomanda(quizzes.Quizzes[numeroDomanda])
		punteggio += risultato
		AggiornaPunteggio(numeroDomanda, risultato)
		scriviEsito(risultato, risposta)
	}
	fmt.Printf("Hai totalizzato %d/30 punti!\n", punteggio)

}
