package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
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
const numeroQuesiti = 166
const numeroEpisodi = 30

var cdf [166]float32
var quizzes Quizzes
var punteggi Punteggi

func GuessMultipleAnswer() {
	rand.Seed(time.Now().UnixNano())

	LoadQuestions(fileDomande, &quizzes)

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

func scriviDomanda(quiz Quiz) (int, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s%s%s\n\n", bold, quiz.Domanda, colorReset)
	for numero, risposta := range quiz.Risposte {
		fmt.Printf("[%s] %s\n", numero, risposta)
	}

	fmt.Print("\n[>] ")
	risposta, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	risposta = strings.TrimSuffix(risposta, "\n")
	if risposta == quiz.Soluzione {
		return 1, ""
	}
	return 0, quiz.Risposte[quiz.Soluzione]
}

func scriviEsito(risultato int, risposta string) {
	if risultato == 1 {
		fmt.Printf("%s[!] Corretto!\n\n%s", colorGreen, colorReset)
	} else {
		fmt.Printf("%s[!] Sbagliato!\n", colorRed)
		fmt.Printf("[>] La risposta corretta era: %s%s%s\n\n", underlined, risposta, colorReset)
	}
	fmt.Printf("%s\n\n", "________________________________________________________________________________")
}

type Punteggi struct {
	Punteggi []Punteggio
}

type Punteggio struct {
	Corrette  float32
	Sbagliate float32
}

const filePunteggi = "/home/ginko/dev/psychopathoquiz/punteggi.json"

func ScegliDomanda() int {
	ComputeCDF()
	r := rand.Float32()

	domanda := 0
	for r > cdf[domanda] {
		domanda++
	}
	return domanda
}

func AggiornaPunteggio(numeroDomanda int, risultato int) {
	loadScore()
	if risultato == 1 {
		punteggi.Punteggi[numeroDomanda].Corrette += 1
	} else {
		punteggi.Punteggi[numeroDomanda].Sbagliate += 1
	}
	byteFile, err := json.Marshal(punteggi)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filePunteggi, byteFile, 0644)
	if err != nil {
		log.Fatal(err)
	}
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

	testCDF()

	for i := range cdf {
		if i > 0 {
			pointer[i] += pointer[i-1]
		}
	}
}

func testCDF() {
	outFile, err := os.Create("/home/ginko/dev/psychopathoquiz/debugCDF.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i, val := range cdf {
		s := fmt.Sprintf("[%d] [%.4f%%] %s\n", i, val, quizzes.Quizzes[i].Domanda)
		_, err = outFile.Write([]byte(s))
		if err != nil {
			log.Fatal(err)
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
