package day12

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Node struct {
	ID       int
	Children []*Node
}

func (n Node) Walk(seen map[*Node]struct{}) {
	for _, child := range n.Children {
		if _, ok := seen[child]; ok {
			continue
		}

		seen[child] = struct{}{}
		child.Walk(seen) // map is a reference type so no need to return
	}

	return
}

func Part1(input io.Reader, startID int) (int, error) {
	nodes := map[int]*Node{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.SplitN(scanner.Text(), " ", 3)
		nodeID, err := strconv.Atoi(line[0])
		if err != nil {
			return 0, err
		}

		if _, ok := nodes[nodeID]; !ok {
			nodes[nodeID] = &Node{ID: nodeID}
		}
		node := nodes[nodeID]

		for _, child := range strings.Split(line[2], ", ") {
			childID, err := strconv.Atoi(child)
			if err != nil {
				return 0, err
			}

			if _, ok := nodes[childID]; !ok {
				nodes[childID] = &Node{ID: childID}
			}

			node.Children = append(node.Children, nodes[childID])
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	if _, ok := nodes[startID]; !ok {
		return 0, fmt.Errorf("node ID not found: %d", startID)
	}

	members := map[*Node]struct{}{}
	nodes[startID].Walk(members)

	return len(members), nil
}
