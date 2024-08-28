package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	level := flag.String("level", "DEBUG", "Log level to filter for")
	flag.Parse()

	f, err := os.Open("./log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bufferedReader := bufio.NewReader(f)

	for line, err := bufferedReader.ReadString('\n'); err == nil; line, err = bufferedReader.ReadString('\n') {
		if strings.Contains(line, *level) {
			fmt.Println(line)
		}
	}
}
