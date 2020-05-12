package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
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

func main() {

	rand.Seed(time.Now().UnixNano())
	jsonData := readJson()
	database := caricaDomande(jsonData)
	fmt.Println(database[2].Domanda)
	fmt.Println(rand.Intn(numeroQuesiti))

	var punteggio int
	CallClear()
	fmt.Printf("[!] Iniziamo!\n\n")
	for i := 1; i <= numeroEpisodi; i++ {

		risultato, risposta := scriviDomanda(database[rand.Intn(numeroQuesiti)])
		punteggio += risultato
		CallClear()
		if risultato == 1 {
			fmt.Printf("[!] Corretto!\n\n")
		} else {
			fmt.Printf("[!] Sbagliato!\n")
			fmt.Printf("[>] La risposta corretta era: %s\n\n", risposta)
		}
	}
	fmt.Printf("Hai totalizzato %d/30 punti!\n", punteggio)

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

func readJson() map[string]json.RawMessage {
	data, _ := ioutil.ReadFile(filename)
	var jsonData map[string]json.RawMessage
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Fatal(err)
	}
	return jsonData
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

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
