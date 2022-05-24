package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Dive() {
	showCommand, showAmount := readInputTxt()
	fmt.Println(showCommand)
	fmt.Println(showAmount)
}

func readInputTxt() ([]string, []int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	command := make([]string, 0)
	amount := make([]int, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		var n int
		var str string
		_, err := fmt.Sscanf(s.Text(), "%s", "%d", &str, &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		command = append(command, str)
		amount = append(amount, n)
	}
	return command, amount
}
