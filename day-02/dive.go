package day02

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Dive() {
  
}

func readInputTxt() {
  f, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  var lines []string

  s := bufio.NewScanner(f)
  s.Split(bufio.ScanLines)
  for s.Scan() {
    lines = append(lines, s.Text())
    
  }

}
