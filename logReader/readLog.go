package logReader

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFileLog() {
	path := flag.String("path", "test.log", "log path")
	contains := flag.String("contains", "carlos", "string  to search")

	fmt.Print("Starting")

	f, err := os.Open(*path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	reader := bufio.NewReader(f)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		if strings.Contains(str, *contains) {
			fmt.Println(str)
		}
	}

}
