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
	phoneRegex := compilePhoneRegex()

	content, err := readFileContent(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	phoneNumbers := findResults(content, phoneRegex)

	for _, number := range phoneNumbers {
		fmt.Println(number)
	}
}

func compilePhoneRegex() *regexp.Regexp {
	return regexp.MustCompile(`\(?\d{3}\)?[-. ]?\d{3}[-. ]?\d{4}`)
}

func findResults(content string, Regex *regexp.Regexp) []string {
	return Regex.FindAllString(content, -1)
}

func readText() {
	path := "files/1689007676028_text.txt"
	textRegex := compileTextRegex()

	content, err := readFileContent(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows := findResults(content, textRegex)

	for _, row := range rows {
		fmt.Println(row)
	}
}

func compileTextRegex() *regexp.Regexp {
	return regexp.MustCompile(`^[аеєиіїоуюяАЕЄИІЇОУЮЯ][а-яА-Я]*[^аеєиіїоуюя\s]$`)
}

func readFileContent(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("не вдалося відкрити файл: %v", err)
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("помилка читання файлу: %v", err)
	}

	return content, nil
}
