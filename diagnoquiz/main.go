package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
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
	fmt.Println(disturbo.Criteri["A"])
	answer := Input()

	score := CompareStrings(answer, FormatString(disturbo.Nome))
	fmt.Println(score)

}

func CompareStrings(answer []string, solution []string) float32 {
	var lenght int
	for i, solWord := range solution {
		for _, ansWord := range answer {

			lenght = len(solWord)
			if lenght < 3 {
				continue
			}

			if strings.Contains(solWord, ansWord) {
				solution[i] = "_"
			} else {
				solution[i] = string(solWord[:lenght-1])
			}

			fmt.Println(solWord, ansWord)
		}
	}
	fmt.Println(solution)

	return 1.3
}

func Input() []string {
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	answerArray := FormatString(answer)
	return answerArray

}

func FormatString(s string) []string {
	s = strings.ToLower(s)
	s = strings.TrimSuffix(s, "\n")
	sArray := strings.Split(s, " ")

	return sArray

}
