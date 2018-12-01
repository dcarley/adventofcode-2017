package day12

import (
	"bufio"
	"fmt"
	"io"
	"sort"
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

func Parse(input io.Reader) (map[int]*Node, error) {
	nodes := map[int]*Node{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.SplitN(scanner.Text(), " ", 3)
		nodeID, err := strconv.Atoi(line[0])
		if err != nil {
			return nodes, err
		}

		if _, ok := nodes[nodeID]; !ok {
			nodes[nodeID] = &Node{ID: nodeID}
		}
		node := nodes[nodeID]

		for _, child := range strings.Split(line[2], ", ") {
			childID, err := strconv.Atoi(child)
			if err != nil {
				return nodes, err
			}

			if _, ok := nodes[childID]; !ok {
				nodes[childID] = &Node{ID: childID}
			}

			node.Children = append(node.Children, nodes[childID])
		}
	}

	return nodes, scanner.Err()
}

func Part1(input io.Reader, startID int) (int, error) {
	nodes, err := Parse(input)
	if err != nil {
		return 0, err
	}

	if _, ok := nodes[startID]; !ok {
		return 0, fmt.Errorf("node ID not found: %d", startID)
	}
	members := map[*Node]struct{}{}
	nodes[startID].Walk(members)

	return len(members), nil
}

func Part2(input io.Reader) (int, error) {
	nodes, err := Parse(input)
	if err != nil {
		return 0, err
	}

	groups := map[string]struct{}{}
	for _, node := range nodes {
		members := map[*Node]struct{}{}
		node.Walk(members)

		group := make([]*Node, 0, len(members))
		for node, _ := range members {
			group = append(group, node)
		}

		sort.Slice(group, func(i, j int) bool {
			return group[i].ID < group[j].ID
		})
		groups[fmt.Sprintf("%v", group)] = struct{}{}
	}

	return len(groups), nil
}
