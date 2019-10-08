// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved."
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Package struct {
	id       string
	start    string
	target   string
	planTime string
	pri      string
}

func main() {

	log.Print("Start...")

	delayFile:= os.Args[1];
	nodeFile := os.Args[2];
	packetFile := os.Args[3]
	answerFile := os.Args[4]

	allPackege := getAllPackage(packetFile)

	// calculate

	output(allPackege, answerFile)

}

func getAllPackage(packetFile string) []Package {
	var allPackege []Package
	inputFile, inputError := os.Open(packetFile)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return nil
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if strings.HasPrefix(inputString, "#") {
			continue
		}
		inputString = strings.Replace(inputString, "(", "", -1)
		inputString = strings.Replace(inputString, ")", "", -1)
		inputString = strings.Replace(inputString, " ", "", -1)
		s := strings.Split(inputString, ",")
		aPackage := Package{}
		aPackage = Package{s[0], s[1], s[2], s[3], s[4]}
		allPackege = append(allPackege, aPackage)
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			break
		}
	}
	return allPackege
}

func output(packages []Package, answer string) {
	fd, _ := os.OpenFile(answer, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	var j int
	for j = 0; j < len(packages); j++ {
		var a string = "(" + packages[j].id + "," + packages[j].planTime + "," + packages[j].start + "," + packages[j].target + ")\n"
		buf := []byte(a)
		fd.Write(buf)
	}

	fd.Close()

}
