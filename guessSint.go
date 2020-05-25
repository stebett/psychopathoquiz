package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GuessSint() {
	rand.Seed(time.Now().UnixNano())
	LoadJson(fileDiagnosi, &diagnosi)

	var totalScore float32
	for i := 0; i < 20; i++ {
		PrintLine()
		totalScore += SintQuiz()
	}

	fmt.Printf("[!] Hai totalizzato %2.1f punti!", totalScore)
}

func SintQuiz() float32 {
	disturbo := LoadSint()
	fmt.Printf("[#] %s%s%s\n", bold, disturbo.Nome, colorReset)
	score := SintEpisode(disturbo)
	return score
}

func SintEpisode(disturbo Disturbo) (totalScore float32) {
	totalScore += AskSintomiCognitivi(disturbo)
	totalScore += AskSintomiComportamentali(disturbo)
	totalScore += AskSintomiEmotivi(disturbo)
	totalScore += AskSintomiNeurovegetativi(disturbo)
	return totalScore
}

func AskSintomiCognitivi(disturbo Disturbo) (partialScore float32) {
	var bestMatch string
	var score float32
	if disturbo.checkCognitivi() {
		fmt.Printf("\n[+] %sSintomi cognitivi%s\n", bold, colorReset)
		sints := CopySpecs(disturbo.Sintomi.Cognitivi)
		for i := len(sints); i > 0; i-- {
			bestMatch, score = compareSpecAnswer(disturbo, sints)
			fmt.Printf("\n[!] %s%s%s\n", underlined, sints[bestMatch], colorReset)
			partialScore += score
			delete(sints, bestMatch)
		}
		fmt.Printf("\n[+] %sSoluzioni:%s\n", bold, colorReset)
		for i, val := range disturbo.Sintomi.Cognitivi {
			_, ok := sints[i]
			if ok {
				fmt.Printf("[%s] %s%s%s\n", i, colorRed, val, colorReset)
			} else {
				fmt.Printf("[%s] %s%s%s\n", i, colorGreen, val, colorReset)
			}
		}
	}
	return partialScore
}

func AskSintomiComportamentali(disturbo Disturbo) (partialScore float32) {
	var bestMatch string
	var score float32
	if disturbo.checkComportamentali() {
		fmt.Printf("\n[+] %sSintomi comportamentali%s\n", bold, colorReset)
		sints := CopySpecs(disturbo.Sintomi.Comportamentali)
		for i := len(sints); i > 0; i-- {
			bestMatch, score = compareSpecAnswer(disturbo, sints)
			fmt.Printf("\n[!] %s%s%s\n", underlined, sints[bestMatch], colorReset)
			partialScore += score
			delete(sints, bestMatch)
		}
		fmt.Printf("\n[+] %sSoluzioni:%s\n", bold, colorReset)
		for i, val := range disturbo.Sintomi.Comportamentali {
			_, ok := sints[i]
			if ok {
				fmt.Printf("[%s] %s%s%s\n", i, colorRed, val, colorReset)
			} else {
				fmt.Printf("[%s] %s%s%s\n", i, colorGreen, val, colorReset)
			}
		}
	}
	return partialScore
}

func AskSintomiNeurovegetativi(disturbo Disturbo) (partialScore float32) {
	var bestMatch string
	var score float32
	if disturbo.checkNeurovegetativi() {
		fmt.Printf("\n[+] %sSintomi cognitivi%s\n", bold, colorReset)
		sints := CopySpecs(disturbo.Sintomi.Neurovegetativi)
		for i := len(sints); i > 0; i-- {
			bestMatch, score = compareSpecAnswer(disturbo, sints)
			fmt.Printf("\n[!] %s%s%s\n", underlined, sints[bestMatch], colorReset)
			partialScore += score
			delete(sints, bestMatch)
		}
		fmt.Printf("\n[+] %sSoluzioni:%s\n", bold, colorReset)
		for i, val := range disturbo.Sintomi.Neurovegetativi {
			_, ok := sints[i]
			if ok {
				fmt.Printf("[%s] %s%s%s\n", i, colorRed, val, colorReset)
			} else {
				fmt.Printf("[%s] %s%s%s\n", i, colorGreen, val, colorReset)
			}
		}
	}
	return partialScore
}

func AskSintomiEmotivi(disturbo Disturbo) (partialScore float32) {
	var bestMatch string
	var score float32
	if disturbo.checkEmotivi() {
		fmt.Printf("\n[+] %sSintomi cognitivi%s\n", bold, colorReset)
		sints := CopySpecs(disturbo.Sintomi.Emotivi)
		for i := len(sints); i > 0; i-- {
			bestMatch, score = compareSpecAnswer(disturbo, sints)
			fmt.Printf("\n[!] %s%s%s\n", underlined, sints[bestMatch], colorReset)
			partialScore += score
			delete(sints, bestMatch)
		}
		fmt.Printf("\n[+] %sSoluzioni:%s\n", bold, colorReset)
		for i, val := range disturbo.Sintomi.Emotivi {
			_, ok := sints[i]
			if ok {
				fmt.Printf("[%s] %s%s%s\n", i, colorRed, val, colorReset)
			} else {
				fmt.Printf("[%s] %s%s%s\n", i, colorGreen, val, colorReset)
			}
		}
	}
	return partialScore
}

func LoadSint() (disturbo Disturbo) {
	n := rand.Intn(numeroDisturbi)
	disturbo = diagnosi.Diagnosi[n]
	if !(disturbo.checkEmotivi() && disturbo.checkCognitivi() && disturbo.checkComportamentali() && disturbo.checkNeurovegetativi()) {
		disturbo = LoadSint()
	}
	return disturbo
}

func (d Disturbo) checkEmotivi() bool {
	if d.Sintomi.Emotivi["1"] == "" {
		return false
	} else {
		return true
	}
}

func (d Disturbo) checkCognitivi() bool {
	if d.Sintomi.Cognitivi["1"] == "" {
		return false
	} else {
		return true
	}
}

func (d Disturbo) checkComportamentali() bool {
	if d.Sintomi.Comportamentali["1"] == "" {
		return false
	} else {
		return true
	}
}
func (d Disturbo) checkNeurovegetativi() bool {
	if d.Sintomi.Neurovegetativi["1"] == "" {
		return false
	} else {
		return true
	}
}