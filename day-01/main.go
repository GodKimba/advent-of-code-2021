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
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		slc = append(slc, n)
	}
	for i := 0; i < len(slc)-1; i++ {
		if slc[i] < slc[i+1] {
			sum += 1
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
