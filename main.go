package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	displayLines          = flag.Bool("n", false, "display lines numbers")
	displayLinesSkipBlank = flag.Bool("b", false, "display lines numbers skipping blank line")

	scanner *bufio.Scanner
)

func main() {

	initializeFlags()

	if hasPipe() {
		scanner = bufio.NewScanner(os.Stdin)
		printLines(scanner)
	} else {
		for _, file := range flag.Args() {
			err := processFile(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error processing file %s: %v\n", file, err)
			}
		}

	}

}

func initializeFlags() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()

		fmt.Fprintf(w, "%s\nA stupid cat(1) clone \n", os.Args[0])

		flag.PrintDefaults()

		fmt.Fprintf(w, "...\n")

	}

	flag.Parse()
}

func hasPipe() bool {
	in := os.Stdin
	stat, err := in.Stat()

	if err != nil {
		panic(err)
	}

	return stat.Mode()&os.ModeNamedPipe != 0
}

func processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	defer file.Close()

	printLines(scanner)

	return nil
}

func printLines(scanner *bufio.Scanner) {
	for line := 1; scanner.Scan(); {
		txt := scanner.Text()

		if *displayLines || *displayLinesSkipBlank {
			if len(txt) == 0 && *displayLinesSkipBlank {
				fmt.Println(txt)
				continue
			}
			fmt.Printf("%d\t%s\n", line, txt)
		} else {
			fmt.Println(txt)
		}

		line++

	}
}
