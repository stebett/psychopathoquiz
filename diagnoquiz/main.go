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

func CompareStrings(answer, solution []string) (score float32) {
	for _, ansWord := range answer {
		var distances []int
		for _, solWord := range solution {
			fmt.Println("Call levensthein with words: ", ansWord, solWord)
			fmt.Println("Levensthein distance: ", Levenshtein(ansWord, solWord))
			distances = append(distances, Levenshtein(ansWord, solWord))
		}
		fmt.Println("distances: ", distances)
		fmt.Println("min distances: ", minArray(distances))
		score += minArray(distances)
	}
	fmt.Println(solution)

	score /= lenStringArray(solution)

	return score
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

func Levenshtein(s, t string) (distance int) {
	//Todo: This is not actually levenshtein distance, need to apply dynamical programming
	if len(s) == 0 {
		return len(t)
	}

	if len(t) == 0 {
		return len(s)
	}

	if t == s {
		return 0
	}

	sRunes := []rune(s)
	tRunes := []rune(t)
	maxLenght := min(len(t), (len(s)))
	for i := 0; i < maxLenght; i++ {
		if sRunes[i] != tRunes[i] {
			distance++
		}

	}
	return distance
}

func argMin(s []int) float32 {
	// Todo: check this works well
	var index int
	for i := range s {
		if s[i] > s[index] {
			continue
		}

		for l := range s {
			if s[i] > s[l] {
				index = l
				continue
			}
		}
	}
	return float32(index)
}

func lenStringArray(s []string) float32 {
	var lenght int
	for _, str := range s {
		lenght += len(str)
	}

	return float32(lenght)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minArray(s []int) float32 {
	tmpMin := s[0]
	for _, val := range s {
		tmpMin = min(tmpMin, val)
	}

	return float32(tmpMin)
}
