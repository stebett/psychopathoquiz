package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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

func Input() []string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n[>] ")
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
func CompareStrings(answer, solution []string) (score float32) {
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
