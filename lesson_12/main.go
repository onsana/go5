package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	readNumbers()
	readText()
}

func readNumbers() {
	const path = "files/1689007675141_numbers.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Помилка відкриття файлу:", err)
		return
	}
	defer file.Close()
	re := regexp.MustCompile(`\(?\d{3}\)?[-. ]?\d{3}[-. ]?\d{4}`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			fmt.Println("Знайдено телефонний номер:", match)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Помилка читання файлу:", err)
	}
}

func readText() {
	path := "files/1689007676028_text.txt"

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Помилка відкриття файлу:", err)
		return
	}
	defer file.Close()

	vowelConsonantRe := regexp.MustCompile(`^[аеєиіїоуюяАЕЄИІЇОУЮЯ][а-яА-Я]*[^аеєиіїоуюя\s]$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		vowelConsonantMatches := vowelConsonantRe.FindAllString(line, -1)
		for _, match := range vowelConsonantMatches {
			fmt.Println("Знайдено слово (голосна-прголосна):", match)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Помилка читання файлу:", err)
	}
}

func hasDoubleLetter(word string) bool {
	runes := []rune(word)
	for i := 0; i < len(runes)-2; i++ {
		if runes[i] == runes[i+2] {
			return true
		}
	}
	return false
}
