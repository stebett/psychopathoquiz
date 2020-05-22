package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const numeroDisturbi = 37

func main() {
	rand.Seed(time.Now().UnixNano())
	loadQuestions()

	QuizCriterioA()
}

func QuizCriterioA() {
	n := rand.Intn(numeroDisturbi)
	disturbo := diagnosi.Diagnosi[n]
	disturbo.PrintCriterio("A")

	answer := Input()
	fmt.Printf("\nPrinting answer ->>%s<<-\n", answer)
	if answer[0] == "" {
		return
	}
	if answer[0] == "cognitivi" {
		disturbo.Sintomi.PrintSintomiCognitivi(0)
	}
	if answer[0] == "comportamentali" {
		disturbo.Sintomi.PrintSintomiComportamentali(0)
	}
	if answer[0] == "emotivi" {
		disturbo.Sintomi.PrintSintomiEmotivi(0)
	}
	if answer[0] == "neurovegetativi" {
		disturbo.Sintomi.PrintSintomiNeurovegetativi(0)
	}

	score := CompareStrings(answer, FormatString(disturbo.Nome))
	ProcessScore(score, disturbo)

}

func (d Disturbo) PrintCriterio(letter string) {
	criterio := d.Criteri[letter]
	fmt.Printf("%s%s%s\n", bold, criterio, colorReset)
	if strings.Contains(criterio, "sintomi") {
		// d.Sintomi.PrintSintomiCognitivi(0)
		// d.Sintomi.PrintSintomiComportamentali(0)
		// d.Sintomi.PrintSintomiEmotivi(0)
		// d.Sintomi.PrintSintomiNeurovegetativi(0)
	}
}

func (s Sintomo) PrintSintomiCognitivi(num int) {
	var index string
	sKind := s.Cognitivi

	fmt.Printf("\n[+] %sSintomi Cognitivi%s\n", bold, colorReset)

	if num == 0 || num > len(sKind) {
		num = len(sKind)
	}
	slice := rand.Perm(num)
	for i := 0; i < len(sKind); i++ {
		index = strconv.Itoa(slice[i] + 1)
		fmt.Printf("[%s] %s\n", index, sKind[index])
	}
}
func (s Sintomo) PrintSintomiComportamentali(num int) {
	var index string
	sKind := s.Comportamentali

	fmt.Printf("\n[+] %sSintomi Comportamentali%s\n", bold, colorReset)

	if num == 0 || num > len(sKind) {
		num = len(sKind)
	}
	slice := rand.Perm(num)
	for i := range slice {
		index = strconv.Itoa(slice[i] + 1)
		fmt.Printf("[%s] %s\n", index, sKind[index])
	}
}
func (s Sintomo) PrintSintomiEmotivi(num int) {
	var index string
	sKind := s.Emotivi

	fmt.Printf("\n[+] %sSintomi Emotivi%s\n", bold, colorReset)

	if num == 0 || num > len(sKind) {
		num = len(sKind)
	}
	slice := rand.Perm(num)
	for i := range slice {
		index = strconv.Itoa(slice[i] + 1)
		fmt.Printf("[%s] %s\n", index, sKind[index])
	}
}
func (s Sintomo) PrintSintomiNeurovegetativi(num int) {
	var index string
	sKind := s.Neurovegetativi

	fmt.Printf("\n[+] %sSintomi Neurovegetativi%s\n", bold, colorReset)

	if num == 0 || num > len(sKind) {
		num = len(sKind)
	}
	slice := rand.Perm(num)
	for i := range slice {
		index = strconv.Itoa(slice[i] + 1)
		fmt.Printf("[%s] %s\n", index, sKind[index])
	}
}

func ProcessScore(score float32, disturbo Disturbo) {
	if score < 0.1 {
		fmt.Printf("%s%s%s\n", colorGreen, "[!] Perfetto!", colorReset)
		return
	} else if score < 0.3 {
		fmt.Printf("%s%s%s\n", colorGreen, "[!] Sembra giusto, controlla la risposta!", colorReset)
	} else if score < 0.5 {
		fmt.Printf("%s%s%s\n", colorRed, "[!] Sembra sbagliato, controlla la risposta!", colorReset)
	} else {
		fmt.Printf("%s%s%s\n", colorRed, "[!] Nope", colorReset)
	}
	fmt.Printf("%s%s%s%s%s\n\n", "[>]", underlined, " La risposta era: ", disturbo.Nome, colorReset)
}

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"

const bold = "\033[1m"
const underlined = "\033[4m"
