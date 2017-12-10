package day07

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Node struct {
	Parent string
	Weight int
}

type Nodes map[string]Node

func Parse(input io.Reader) (Nodes, string, error) {
	nodes := Nodes{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.SplitN(scanner.Text(), " ", 4)
		name := line[0]

		weight, err := strconv.Atoi(strings.Trim(line[1], "()"))
		if err != nil {
			return Nodes{}, "", err
		}

		if node, exists := nodes[name]; exists {
			node.Weight = weight
			nodes[name] = node
		} else {
			nodes[name] = Node{
				Weight: weight,
			}
		}

		if len(line) > 2 {
			for _, child := range strings.Split(line[3], ", ") {
				if node, exists := nodes[child]; exists {
					node.Parent = name
					nodes[child] = node
				} else {
					nodes[child] = Node{
						Parent: name,
					}
				}
			}
		}
	}

	var root string
	for name, node := range nodes {
		if node.Parent == "" {
			root = name
			break
		}
	}

	return nodes, root, scanner.Err()
}

type UnbalancedNode struct {
	Name   string
	Weight int
}

func TraverseChildren(nodes Nodes, parent string) (int, *UnbalancedNode) {
	branchWeights := map[int][]string{}

	for name, node := range nodes {
		if node.Parent == parent {
			branchWeight, unbalanced := TraverseChildren(nodes, name)
			if unbalanced != nil {
				return 0, unbalanced
			}

			branchWeight += node.Weight
			branchWeights[branchWeight] = append(branchWeights[branchWeight], name)
		}
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
		name := branchWeights[incorrectWeight][0]
		correction := incorrectWeight - correctWeight

		return 0, &UnbalancedNode{
			Name:   name,
			Weight: nodes[name].Weight - correction,
		}
	}

	return sum, nil
}

func Part1(input io.Reader) (string, error) {
	_, root, err := Parse(input)
	if err != nil {
		return "", err
	}

	return root, nil
}

func Part2(input io.Reader) (string, int, error) {
	nodes, root, err := Parse(input)
	if err != nil {
		return "", 0, err
	}

	_, unbalanced := TraverseChildren(nodes, root)

	return unbalanced.Name, unbalanced.Weight, nil
}
