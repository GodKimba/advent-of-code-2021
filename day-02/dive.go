package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var n int
	var str string

	type steps []submarineInstrucs

	step := steps{}
	for s.Scan() {
		_, err := fmt.Sscanf(s.Text(), "%s %d", &str, &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		step = append(step, submarineInstrucs{str, n}, submarineInstrucs{str, n})
	}
	fmt.Println(step)

}

type submarineInstrucs struct {
	commands string
	amount   int
}
