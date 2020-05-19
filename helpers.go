package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func ScegliDomanda() int {
	CalcolaProbCumulativa()
	r := rand.Float32()

	domanda := 0
	for r > cdf[domanda] {
		domanda++
	}
	return domanda
}

func CalcolaProbCumulativa() {
	punteggi := CaricaPunteggi(numeroDomanda, filename)

}

type Punteggio struct {
	Corrette  int
	Sbagliate int
}

func AggiornaPunteggio(numeroDomanda int, risultato int, filename string) {
	return

}

func CaricaPunteggi(numeroDomanda int, filename string) []*Punteggio {
	jsonData := readJson(filename)

	var punteggi []*Punteggio

	for i := 1; i <= numeroQuesiti; i++ {
		quesito := "Quesito" + strconv.Itoa(i)
		punteggi = append(punteggi, CaricaUnPunteggio(quesito, jsonData))
	}

	return punteggi
}

func CaricaUnPunteggio(numeroQuesito string, jsonData map[string]json.RawMessage) *Punteggio {

	punteggio := new(Punteggio)
	var quesito map[string]json.RawMessage

	err := json.Unmarshal(jsonData[numeroQuesito], &quesito)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(quesito["Corrette"], &punteggio.Corrette)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(quesito["Sbagliate"], &punteggio.Sbagliate)
	if err != nil {
		log.Fatal(err)
	}

	return punteggio
}

func scriviEsito(risultato int, risposta string) {
	if risultato == 1 {
		fmt.Printf("[!] Corretto!\n\n")
	} else {
		fmt.Printf("[!] Sbagliato!\n")
		fmt.Printf("[>] La risposta corretta era: %s\n\n", risposta)
	}
}

func caricaDomande(jsonData map[string]json.RawMessage) []*Quiz {
	var database []*Quiz

	for i := 1; i <= numeroQuesiti; i++ {
		quesito := "Quesito" + strconv.Itoa(i)
		database = append(database, caricaUnaDomanda(quesito, jsonData))
	}

	return database
}

func caricaUnaDomanda(numeroQuesito string, jsonData map[string]json.RawMessage) *Quiz {
	quiz := new(Quiz)
	var quesito map[string]json.RawMessage

	err := json.Unmarshal(jsonData[numeroQuesito], &quesito)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(quesito["Domanda"], &quiz.Domanda)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(quesito["Risposte"], &quiz.Risposte)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(quesito["Soluzione"], &quiz.Soluzione)
	if err != nil {
		log.Fatal(err)
	}

	return quiz
}

func scriviDomanda(quiz *Quiz) (int, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(quiz.Domanda)
	fmt.Println()
	for numero, risposta := range quiz.Risposte {
		fmt.Printf("[%s] %s\n", numero, risposta)
	}

	fmt.Print("\n[*] Scrivi il numero della risposta: ")
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

func readJson(filename string) map[string]json.RawMessage {
	data, _ := ioutil.ReadFile(filename)
	var jsonData map[string]json.RawMessage
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Fatal(err)
	}
	return jsonData
}
