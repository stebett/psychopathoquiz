package main

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
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
	fmt.Printf("[#] %s\n", disturbo.Nome)
	score := SpecEpisode(disturbo)
	return score
}

func SpecEpisode(disturbo Disturbo) (totalScore float32) {
	var bestMatch string
	var score float32

	specs := CopySpecs(disturbo.Specificatori)

	for i := len(specs); i > 0; i-- {
		// fmt.Printf("[*] Mancano %d specificatori!\n", i)
		bestMatch, score = compareSpecAnswer(disturbo, specs)
		fmt.Printf("\n[!] %s\n", specs[bestMatch])
		totalScore += score
		delete(specs, bestMatch)
	}

	fmt.Printf("\n[+] Soluzioni:\n")
	for i, val := range disturbo.Specificatori {
		_, ok := specs[i]
		if ok {
			fmt.Printf("[%s] %s\n", i, val)
		} else {
			fmt.Printf("[%s] %s\n", i, val)
		}
	}
	return totalScore
}

func compareSpecAnswer(disturbo Disturbo, specs map[string]string) (index string, score float32) {
	bestScore := float32(1)

	answer := Input()
	for i, val := range specs {
		score = CompareStrings(answer, FormatString(val))
		if score < bestScore {
			bestScore = score
			index = i
		}
	}
	return index, bestScore
}

func CopySpecs(oldSpecs map[string]string) map[string]string {
	specs := make(map[string]string)
	for i, val := range oldSpecs {
		specs[i] = removeParenthesis(val)
	}
	return specs
}

func LoadSpec() (disturbo Disturbo) {
	n := rand.Intn(numeroDisturbi)
	disturbo = diagnosi.Diagnosi[n]
	if disturbo.Specificatori["1"] == "" {
		disturbo = LoadSpec()
	}
	return disturbo
}

func removeParenthesis(s string) string {
	reg, err := regexp.Compile(`\((.*)\)`)
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(s, "")
	return processedString
}
