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
	
	sum := 0
	depth := slc[0] + slc[1] + slc[2]

	for i := 0; i < len(slc) - 2; i++ {
		if  slc[i] + slc[i+1] + slc[i+2] > depth{
			sum += 1
		}
		depth = slc[i] + slc[i+1] + slc[i+2]
	}
	fmt.Println(sum)
	
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
