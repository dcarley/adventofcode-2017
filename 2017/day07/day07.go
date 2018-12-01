package day07

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Node struct {
	Name     string
	Weight   int
	Parent   *Node
	Children []*Node
}

func (n Node) Balance() (int, *Node) {
	branchWeights := map[int][]*Node{}

	for _, node := range n.Children {
		branchWeight, unbalanced := node.Balance()
		if unbalanced != nil {
			return 0, unbalanced
		}

		branchWeight += node.Weight
		branchWeights[branchWeight] = append(branchWeights[branchWeight], node)
	}

	var sum, correctWeight, incorrectWeight int

	for weight, names := range branchWeights {
		sum += weight * len(names)

		if len(names) == 1 {
			incorrectWeight = weight
		} else {
			correctWeight = weight
		}
	}

	if incorrectWeight != 0 {
		unbalanced := branchWeights[incorrectWeight][0]
		unbalanced.Weight -= (incorrectWeight - correctWeight)

		return 0, unbalanced
	}

	return sum, nil
}

func Parse(input io.Reader) (*Node, error) {
	nodes := map[string]*Node{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.SplitN(scanner.Text(), " ", 4)
		name := line[0]

		weight, err := strconv.Atoi(strings.Trim(line[1], "()"))
		if err != nil {
			return nil, err
		}

		if _, ok := nodes[name]; !ok {
			nodes[name] = &Node{Name: name}
		}
		node := nodes[name]
		node.Weight = weight

		if len(line) > 2 {
			for _, childName := range strings.Split(line[3], ", ") {
				if _, ok := nodes[childName]; !ok {
					nodes[childName] = &Node{Name: childName}
				}

				child := nodes[childName]
				child.Parent = node
				node.Children = append(node.Children, child)
			}
		}
	}

	var root *Node
	for _, node := range nodes {
		if node.Parent == nil {
			root = node
			break
		}
	}

	return root, scanner.Err()
}

func Part1(input io.Reader) (string, error) {
	root, err := Parse(input)
	if err != nil {
		return "", err
	}

	return root.Name, nil
}

func Part2(input io.Reader) (string, int, error) {
	root, err := Parse(input)
	if err != nil {
		return "", 0, err
	}

	_, unbalanced := root.Balance()

	return unbalanced.Name, unbalanced.Weight, nil
}
