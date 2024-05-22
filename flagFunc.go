package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	//fmt.Printf("	%d %s\n", byteCount, fileName)
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

	//fmt.Printf("	%d %s\n", lineCount, fileName)

	return lineCount, nil
}

//func numWords(fileName string) (int, error) {
//
//	file, err := os.Open(fileName)
//	if err != nil {
//		return -1, err
//	}
//	defer file.Close()
//
//	//data, _ := io.ReadAll(file)
//	//wordCount := len(strings.Split(string(data), " "))
//	wordCount := 0
//	wordScanner := bufio.NewScanner(file)
//	for wordScanner.Scan() {
//		text := strings.Split(wordScanner.Text(), " ")
//		if len(text) == 0 {
//			continue
//		}
//		for _, i := range text {
//			if i != " " {
//				wordCount += 1
//			}
//		}
//	}
//	fmt.Printf("	%d %s\n", wordCount, fileName)
//	return wordCount, nil
//}

func numWords(fileName string) (int, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	wordCount := 0
	wordScanner := bufio.NewScanner(file)
	wordScanner.Split(bufio.ScanWords) // Use bufio.ScanWords for word splitting

	for wordScanner.Scan() {
		wordCount++
	}
	//fmt.Printf("	%d %s\n", wordCount, fileName)
	return wordCount, nil
}

func numChar(fileName string) (int, error) {
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
	//fmt.Printf("	%d %s\n", charCount, fileName)
	return charCount, nil
}

func defaultFlag(fileName string) {
	bytes, err := numBytes(fileName)
	words, err := numWords(fileName)
	lines, err := numLines(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Printf("	%v %v %v %v\n", bytes, words, lines, fileName)

}
