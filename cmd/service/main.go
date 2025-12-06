package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fathimasithara01/realimage-challenge-2016/internal/permissions"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: service <permissions.txt> <queries.txt>")
	}

	permFile := os.Args[1]
	queryFile := os.Args[2]

	svc := permissions.NewService()

	loadPermissions(svc, permFile)
	ch := permissions.NewChecker(svc)

	runQueries(ch, queryFile)
}

func loadPermissions(svc *permissions.Service, file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("failed to read permissions file: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var currentDistributor string

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}

		if !strings.Contains(line, ":") {
			currentDistributor = line
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		switch strings.ToUpper(key) {
		case "INCLUDE":
			svc.AddRule(currentDistributor, permissions.Include, val)
		case "EXCLUDE":
			svc.AddRule(currentDistributor, permissions.Exclude, val)
		}
	}
}

func runQueries(checker *permissions.Checker, file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("failed to read queries file: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}

		ok := checker.Evaluate(permissions.Line{
			Distributor: parts[0],
			Region:      parts[1],
		})

		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
