package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("camel.in")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	answer := 1
	line = strings.TrimSpace(line)
	for _, ch := range line {
		str := string(ch)
		if strings.ToUpper(str) == str {
			answer++
			fmt.Println(str)
		}
	}

	fmt.Println(answer)
}
