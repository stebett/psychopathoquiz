package main

import (
	"fmt"
	"math/rand"
	"strconv"
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

	score := AnswerHandler(disturbo)
	ProcessScore(score, disturbo)

}

func AnswerHandler(disturbo Disturbo) (score float32) {
	answer := Input()
	if answer[0] == "" || answer[0] == "help" {
		fmt.Println("Puoi chiedere aiuto scrivendo il numero o la parola associati ai sintomi o criteri che ti servono!")
		fmt.Println("[1] cognitivi")
		fmt.Println("[2] comportamentali")
		fmt.Println("[3] emotivi")
		fmt.Println("[4] neurovegetativi")
		fmt.Println("[5] tempo")
		fmt.Println("[6] B")
		fmt.Println("[7] C")
		fmt.Println("[8] D")
		fmt.Println("[9] E")
		fmt.Println("[10] F")
		return AnswerHandler(disturbo)
	}
	if answer[0] == "cognitivi" || answer[0] == "1" {
		disturbo.Sintomi.PrintSintomiCognitivi(0)
		return AnswerHandler(disturbo)
	}
	if answer[0] == "comportamentali" || answer[0] == "2" {
		disturbo.Sintomi.PrintSintomiComportamentali(0)
		return AnswerHandler(disturbo)
	}
	if answer[0] == "emotivi" || answer[0] == "3" {
		disturbo.Sintomi.PrintSintomiEmotivi(0)
		return AnswerHandler(disturbo)
	}
	if answer[0] == "neurovegetativi" || answer[0] == "4" {
		disturbo.Sintomi.PrintSintomiNeurovegetativi(0)
		return AnswerHandler(disturbo)
	}
	if answer[0] == "tempo" || answer[0] == "5" {
		disturbo.PrintCriterio("Tempo")
		return AnswerHandler(disturbo)
	}
	if answer[0] == "b" || answer[0] == "6" {
		disturbo.PrintCriterio("b")
		return AnswerHandler(disturbo)
	}
	if answer[0] == "c" || answer[0] == "7" {
		disturbo.PrintCriterio("c")
		return AnswerHandler(disturbo)
	}
	if answer[0] == "d" || answer[0] == "8" {
		disturbo.PrintCriterio("d")
		return AnswerHandler(disturbo)
	}
	if answer[0] == "e" || answer[0] == "9" {
		disturbo.PrintCriterio("e")
		return AnswerHandler(disturbo)
	}
	if answer[0] == "f" || answer[0] == "10" {
		disturbo.PrintCriterio("f")
		return AnswerHandler(disturbo)
	}

	score = CompareStrings(answer, FormatString(disturbo.Nome))
	return score

}

func (d Disturbo) PrintCriterio(letter string) {
	criterio := d.Criteri[letter]
	fmt.Printf("%s%s%s\n", bold, criterio, colorReset)
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
