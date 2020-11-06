package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	inPath := flag.String("in", ".env", "path/to/.env")
	outPath := flag.String("out", ".env.example", "output path")

	flag.Parse()

	file, err := os.Open(*inPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	out, err := os.Create(*outPath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// Handle empty lines
			_, err := out.WriteString("\n")
			if err != nil {
				panic(err)
			}
		} else if strings.HasPrefix(strings.Trim(line, " "), "#") {
			// Handle comments
			_, err := out.WriteString(fmt.Sprintf("%s\n", line))
			if err != nil {
				panic(err)
			}

		} else {
			key := strings.Split(line, "=")[0]

			_, err := out.WriteString(fmt.Sprintf("%s=\n", key))
			if err != nil {
				panic(err)
			}
		}
	}
}
