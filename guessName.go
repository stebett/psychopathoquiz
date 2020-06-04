package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Diagnosi struct {
	Diagnosi []Disturbo
}

type Disturbo struct {
	Nome                  string
	Criteri               map[string]string
	Sintomi               Sintomo
	Specificatori         map[string]string
	DiagnosiDifferenziale map[string]string
	Prevalenza            string
	Incidenza             string
	DifferenzaGenere      string
	DifferenzaEta         string
	Decorsi               map[string]string
	Prognosi              string
}

type Sintomo struct {
	Emotivi         map[string]string
	Cognitivi       map[string]string
	Neurovegetativi map[string]string
	Comportamentali map[string]string
}

const numeroDisturbi = 37
const fileDiagnosi = "/home/ginko/dev/go/src/github.com/ginkobab/psychopathoquiz/diagnosi.json"

var diagnosi Diagnosi

func GuessName() {
	rand.Seed(time.Now().UnixNano())
	LoadJson(fileDiagnosi, &diagnosi)

	var totalScore float32
	for i := 0; i < 30; i++ {
		PrintLine()
		totalScore += NameQuiz()
	}

	fmt.Printf("[!] Hai totalizzato %2.1f punti!", totalScore)
}

func NameQuiz() float32 {
	n := rand.Intn(numeroDisturbi)
	disturbo := diagnosi.Diagnosi[n]

	disturbo.PrintCriterio("A")

	score := AnswerHandler(disturbo)
	return ProcessScore(score, disturbo)

}

func (d Disturbo) PrintCriterio(letter string) {
	criterio := d.Criteri[letter]
	if criterio != "" {
		fmt.Printf("[%s] %s%s%s\n\n", letter, bold, criterio, colorReset)
	}
}

func (d Disturbo) PrintDemographics() {
	if d.Prevalenza != "" {
		fmt.Printf("[P] %s\n", d.Prevalenza)
	}
	if d.Incidenza != "" {
		fmt.Printf("[I] %s\n", d.Incidenza)
	}
	if d.DifferenzaGenere != "" {
		fmt.Printf("[G] %s\n", d.DifferenzaGenere)
	}
	if d.DifferenzaEta != "" {
		fmt.Printf("[E] %s\n", d.DifferenzaEta)
	}
}

func (d Disturbo) PrintSpecificatori() {
	if d.Specificatori["1"] != "" {
		fmt.Printf("\n[+] %sSpecificatori%s\n", bold, colorReset)
	}
	for i, val := range d.Specificatori {
		fmt.Printf("[%s] %s\n", i, val)
	}
}

func (s Sintomo) PrintSintomiCognitivi(num int) {
	var index string
	sKind := s.Cognitivi

	if sKind["1"] != "" {
		fmt.Printf("\n[+] %sSintomi Cognitivi%s\n", bold, colorReset)
	}

	if num == 0 || num > len(sKind) {
		num = len(sKind)
	}
	slice := rand.Perm(num)
	for i := 0; i < len(sKind); i++ {
		index = strconv.Itoa(slice[i] + 1)
		if sKind[index] != "" {
			fmt.Printf("\t[*] %s\n", sKind[index])
		}
	}
}

func (s Sintomo) PrintSintomiComportamentali(num int) {
	var index string
	sKind := s.Comportamentali

	if sKind["1"] != "" {
		fmt.Printf("\n[+] %sSintomi Comportamentali%s\n", bold, colorReset)
	}

	if num == 0 || num > len(sKind) {
		num = len(sKind)
	}
	slice := rand.Perm(num)
	for i := range slice {
		index = strconv.Itoa(slice[i] + 1)
		if sKind[index] != "" {
			fmt.Printf("\t[*] %s\n", sKind[index])
		}
	}
}

func (s Sintomo) PrintSintomiEmotivi(num int) {
	var index string
	sKind := s.Emotivi

	if sKind["1"] != "" {
		fmt.Printf("\n[+] %sSintomi Emotivi%s\n", bold, colorReset)
	}

	if num == 0 || num > len(sKind) {
		num = len(sKind)
	}
	slice := rand.Perm(num)
	for i := range slice {
		index = strconv.Itoa(slice[i] + 1)
		if sKind[index] != "" {
			fmt.Printf("\t[*] %s\n", sKind[index])
		}
	}
}

func (s Sintomo) PrintSintomiNeurovegetativi(num int) {
	var index string
	sKind := s.Neurovegetativi

	if sKind["1"] != "" {
		fmt.Printf("\n[+] %sSintomi Neurovegetativi%s\n", bold, colorReset)
	}

	if num == 0 || num > len(sKind) {
		num = len(sKind)
	}
	slice := rand.Perm(num)
	for i := range slice {
		index = strconv.Itoa(slice[i] + 1)
		if sKind[index] != "" {
			fmt.Printf("\t[*] %s\n", sKind[index])
		}
	}
}

func (d Disturbo) PrintAll() {
	fmt.Printf("______\n\n")
	d.PrintCriterio("B")
	d.PrintCriterio("C")
	d.PrintCriterio("D")
	d.PrintCriterio("E")
	d.PrintCriterio("F")
	d.PrintCriterio("G")
	d.PrintCriterio("Tempo")
	d.Sintomi.PrintSintomiCognitivi(0)
	d.Sintomi.PrintSintomiEmotivi(0)
	d.Sintomi.PrintSintomiNeurovegetativi(0)
	d.Sintomi.PrintSintomiComportamentali(0)
	fmt.Println()
	d.PrintDemographics()
	d.PrintSpecificatori()
	fmt.Println("____")
}

func AnswerHandler(disturbo Disturbo) (score float32) {
	answer := Input()
	if answer[0] == "" || answer[0] == "help" {
		fmt.Println("[?] Puoi chiedere aiuto scrivendo il numero o la parola associati ai sintomi o criteri che ti servono!")
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
		fmt.Println("[11] tutto")
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
	if answer[0] == "tutto" || answer[0] == "11" {
		disturbo.PrintAll()
		return AnswerHandler(disturbo)
	}

	score = CompareStrings(answer, FormatString(disturbo.Nome))
	return score

}

func ProcessScore(score float32, disturbo Disturbo) float32 {
	if score < 0.1 {
		fmt.Printf("%s%s%s\n", colorGreen, "[!] Perfetto!", colorReset)
		return 1
	} else if score < 0.3 {
		fmt.Printf("%s%s%s\n", colorGreen, "[!] Sembra giusto, controlla la risposta!", colorReset)
		fmt.Printf("%s%s%s%s%s\n\n", "[>]", underlined, " La risposta era: ", disturbo.Nome, colorReset)
		return 0.8
	} else if score < 0.5 {
		fmt.Printf("%s%s%s\n", colorRed, "[!] Sembra sbagliato, controlla la risposta!", colorReset)
		fmt.Printf("%s%s%s%s%s\n\n", "[>]", underlined, " La risposta era: ", disturbo.Nome, colorReset)
		return 0.1
	} else {
		fmt.Printf("%s%s%s\n", colorRed, "[!] Nope", colorReset)
		fmt.Printf("%s%s%s%s%s\n\n", "[>]", underlined, " La risposta era: ", disturbo.Nome, colorReset)
		return 0
	}
}

func Levenshtein(s, t string) (result int) {
	if s == t {
		return 0
	}
	if s == "" {
		return len(t)
	}
	if t == "" {
		return len(s)
	}

	matrix := GenerateMatrix(len(s), len(t))
	if s[0] != t[0] {
		for i := 1; i < len(s); i++ {
			matrix[i][0] += 1
		}
		for i := 1; i < len(t); i++ {
			matrix[0][i] += 1
		}
	}

	for r := 1; r < len(s); r++ {
		for c := 1; c < len(t); c++ {
			if s[r] == t[c] {
				matrix[r][c] = matrix[r-1][c-1]
			} else {
				matrix[r][c] = min(matrix[r][c-1], matrix[r-1][c-1], matrix[r-1][c])
				matrix[r][c]++
			}
		}
	}

	if matrix[len(s)-1][len(t)-1] > len(t) {
		return len(t)
	} else {
		return matrix[len(s)-1][len(t)-1]
	}
}

func GenerateMatrix(rows, cols int) [][]int {

	matrix := make([][]int, rows)

	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		matrix[i][0] = i
	}
	for i := 0; i < cols; i++ {
		matrix[0][i] = i
	}
	return matrix
}

func CompareStrings(answer, solution []string) (score float32) {
	for i := len(answer); i <= len(solution); i++ {
		answer = append(answer, "")
	}
	for _, solWord := range solution {
		var distances []int
		for _, ansWord := range answer {
			distances = append(distances, Levenshtein(ansWord, solWord))
		}
		score += minArray(distances)
	}

	score /= lenStringArray(solution)

	return score
}

func min(args ...int) int {
	m := args[0]
	for _, val := range args {
		if val < m {
			m = val
		}
	}
	return m
}

func lenStringArray(s []string) float32 {
	var lenght int
	for _, str := range s {
		lenght += len(str)
	}

	return float32(lenght)
}

func minArray(s []int) float32 {
	tmpMin := s[0]
	for _, val := range s {
		tmpMin = min(tmpMin, val)
	}

	return float32(tmpMin)
}

func argMin(s []float32) (index int) {
	oldVal := s[0]
	for i, val := range s {
		if val < oldVal {
			index = i
		}
	}
	return index
}
