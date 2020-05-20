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

	fmt.Println(ScegliDomanda())
	var punteggio int
	for i := 1; i <= numeroEpisodi; i++ {

		risultato, risposta := scriviDomanda(quizzes.Quizzes[1])
		punteggio += risultato
		AggiornaPunteggio(1, risultato)
		scriviEsito(risultato, risposta)
	}
	fmt.Printf("Hai totalizzato %d/30 punti!\n", punteggio)

}
