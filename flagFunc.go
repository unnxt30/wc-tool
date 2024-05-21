package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func numBytes(fileName string) (int, error) {
	var byteCount int
	file, err := os.Open(fileName)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	data, _ := io.ReadAll(file)
	byteCount = len(data)
	fmt.Printf("	%d %s\n", byteCount, fileName)
	return byteCount, nil
}

func numLines(fileName string) (int, error) {
	var lineCount int = 0
	file, err := os.Open(fileName)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	lineScanner := bufio.NewScanner(file)
	for lineScanner.Scan() {
		lineCount += 1
	}

	fmt.Printf("	%d %s\n", lineCount, fileName)

	return lineCount, nil
}

func numWords(fileName string) (int, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	data, _ := io.ReadAll(file)
	wordCount := len(strings.Split(string(data), " "))

	fmt.Printf("	%d %s\n", wordCount, fileName)
	return wordCount, nil
}

func numSpecialChar(fileName string) (int, error) {
	var charCount int
	file, err := os.Open(fileName)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	data, _ := io.ReadAll(file)
	for i := range data {
		charCount = i
	}
	charCount += 1
	fmt.Printf("	%d %s\n", charCount, fileName)
	return charCount, nil
}
