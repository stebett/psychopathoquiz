package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func scriviDomanda(quiz Quiz) (int, string) {
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

func scriviEsito(risultato int, risposta string) {
	if risultato == 1 {
		fmt.Printf("[!] Corretto!\n\n")
	} else {
		fmt.Printf("[!] Sbagliato!\n")
		fmt.Printf("[>] La risposta corretta era: %s\n\n", risposta)
	}
}
