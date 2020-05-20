package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"

const bold = "\033[1m"
const underlined = "\033[4m"

func scriviDomanda(quiz Quiz) (int, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s%s%s\n", bold, quiz.Domanda, colorReset)
	fmt.Println()
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
		fmt.Printf("[>] La risposta corretta era: %s%s\n\n%s", underlined, risposta, colorReset)
	}
	fmt.Printf("%s\n\n", "________________________________________________________________________________")
}
