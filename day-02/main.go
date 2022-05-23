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

	sum := 0
	slc := make([]int, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		var n int
		_, err := fmt.Sscan(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %d", s.Text(), err)
		}
		slc = append(slc, n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
