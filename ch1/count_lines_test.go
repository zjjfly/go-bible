package ch1

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestCountLines(t *testing.T) {
	count := make(map[string]int)
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}
	for line, n := range count {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func TestCountLines2(t *testing.T) {
	count := make(map[string]int)
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	for _, line := range strings.Split(string(data), "\n") {
		count[line]++
	}
	for line, n := range count {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
