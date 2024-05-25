package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Row struct {
	Id int
}

func main() {
	text := getAllText()
	word := "wisdom"
	rows := wordSearch(text, word)
	fmt.Println(rows)
	words := textToWords(text)
	frequency := wordFrequency(words)
	fmt.Println(frequency)
}

func getAllText() []string {
	myfile, err := os.Open("text.txt")
	text := []string{}
	if err != nil {
		fmt.Println("Error opening file:", err)
		return text
	}
	defer myfile.Close()

	scanner := bufio.NewScanner(myfile)

	for scanner.Scan() {
		line := scanner.Text()
		text = append(text, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
	return text
}

func textToWords(text []string) []string {
	wordsSlice := []string{}
	for _, v := range text {
		w := strings.Split(v, " ")
		wordsSlice = append(wordsSlice, w...)

	}
	return wordsSlice
}

func wordFrequency(words []string) map[string]int {
	frequency := make(map[string]int)
	for _, v := range words {
		frequency[v]++
	}
	return frequency
}

func wordSearch(text []string, word string) []Row {
	result := []Row{}
	for i, v := range text {
		if strings.Contains(v, word) {
			result = append(result, Row{Id: i})
		}
	}
	return result
}
