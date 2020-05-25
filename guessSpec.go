package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GuessSpec() {
	rand.Seed(time.Now().UnixNano())
	LoadJson(fileDiagnosi, &diagnosi)

	var totalScore float32
	for i := 0; i < 20; i++ {
		PrintLine()
		totalScore += SpecQuiz()
	}

	fmt.Printf("[!] Hai totalizzato %2.1f punti!", totalScore)

}

func SpecQuiz() float32 {
	disturbo := LoadSpec()
	fmt.Printf("[#] %s%s%s\n", bold, disturbo.Nome, colorReset)
	score := SpecEpisode(disturbo)
	return score
}

func SpecEpisode(disturbo Disturbo) (totalScore float32) {
	var bestMatch string
	var score float32

	for i := len(disturbo.Specificatori); i > 0; i-- {
		fmt.Printf("[*] Mancano %d specificatori!\n", i)
		bestMatch, score = compareSpecAnswer(disturbo)
		fmt.Printf("\n[!] Eri vicino a: %s%s%s con un punteggio di %.2f\n", colorGreen, disturbo.Specificatori[bestMatch], colorReset, score)
		totalScore += score
	}
	disturbo.PrintSpecificatori()

	return totalScore
}

func compareSpecAnswer(disturbo Disturbo) (index string, score float32) {
	bestScore := float32(1)

	answer := Input()
	for i, val := range disturbo.Specificatori {
		score = CompareStrings(answer, FormatString(val))
		if score < bestScore {
			bestScore = score
			index = i
		}
	}

	return index, bestScore
}

func LoadSpec() (disturbo Disturbo) {
	n := rand.Intn(numeroDisturbi)
	disturbo = diagnosi.Diagnosi[n]
	if disturbo.Specificatori["1"] == "" {
		disturbo = LoadSpec()
	}
	return disturbo
}
