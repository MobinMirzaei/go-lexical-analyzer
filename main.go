package main

import (
	"fmt"
	"github.com/mobinmirzaei/lexical-analyzer/lexer"
	"os"
)

func main() {

	// read from code.txt and write tokens to result.txt
	inputFileName := "code.txt"
	outputFileName := "result.txt"

	content, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Printf("error to read file. %v\n", err)
		return
	}

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Printf("error to create file. %v\n", err)
		return
	}
	defer outputFile.Close() 


	l := lexer.New(string(content))

	header := fmt.Sprintf("%-20s | %-20s | %-5s\n", "TOKEN TYPE", "LITERAL", "LINE")
	separator := "------------------------------------------------------------\n"
	
	fmt.Print(header, separator)
	outputFile.WriteString(header + separator)

	for {
		tok := l.NextToken()

		lineResult := fmt.Sprintf("%-20s | %-20s | %-5d\n", tok.Type, tok.Literal, tok.Line)

		fmt.Print(lineResult)

		outputFile.WriteString(lineResult)

		if tok.Type == "EOF" {
			break
		}
	}

	fmt.Println("Lexical analysis completed. Results have been saved to result.txt.")

}