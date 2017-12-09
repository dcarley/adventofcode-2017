package day07

import (
	"bufio"
	"io"
	"strings"
)

func Part1(input io.Reader) (string, error) {
	nodes := map[string]string{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.SplitN(scanner.Text(), " ", 4)
		name := line[0]
		if _, exists := nodes[name]; !exists {
			nodes[name] = ""
		}

		if len(line) > 2 {
			for _, child := range strings.Split(line[3], ", ") {
				nodes[child] = name
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	var root string
	for name, parent := range nodes {
		if parent == "" {
			root = name
			break
		}
	}

	return root, nil
}
