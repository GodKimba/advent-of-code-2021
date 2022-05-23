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
	slice := make([]int, 10)
	s := bufio.NewScanner(f)
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		slice = append(slice, n)
	}
	for i := 0; i < len(slice[11:])-1; i++ {
		if slice[i] < slice[i+1] {
			sum += 1
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
	fmt.Println(len(slice) - 10)
}
