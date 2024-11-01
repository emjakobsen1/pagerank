package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*Usage: go run main.go <filename> */
func main() {
	fmt.Printf("LINEAR ALGEBRA 2024 \n")

	/* Command line arguments for file I/O. */
	var filename string

	if len(os.Args) < 2 {
		filename = "p2p-Gnutella08-mod.txt"
	} else {
		filename = os.Args[1]
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	fmt.Printf("File:	%s \n", filename)

	/* Construct the graph and reversed graph. It is an Adjancy List representation of type map int -> int list,
	with keys being the nodes and values being a list of edges associated with the specific entry. */
	graph := make(map[int][]int)

	scanner := bufio.NewScanner(file)

	edgeCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) != 2 {
			log.Fatalf("Invalid line format: %s", line)
		}

		node1, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatalf("Invalid node value: %s", fields[0])
		}
		node2, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("Invalid node value: %s", fields[1])
		}
		graph[node1] = append(graph[node1], node2)
		edgeCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	addDanglingNodes(graph)

	nodesCount := len(graph)

	fmt.Printf("Graph order: %d nodes, size: %d edges\n\n", nodesCount, edgeCount)

	/* Used to check the time running the algorithms. */
	start := time.Now().UnixNano()

	/* Sets the graph, damping factor and how many steps the walk is for the random surf. */
	randomSurf(graph, 0.15, 100000000)

	duration := float64(time.Now().UnixNano()-start) / 1e6
	fmt.Printf("Time:	%.3f milliseconds\n", duration)

	pageRank(graph, nodesCount, edgeCount)

	reversedGraph := make(map[int][]int)
	for node, adjNodes := range graph {
		for _, adjNode := range adjNodes {
			reversedGraph[adjNode] = append(reversedGraph[adjNode], node)
		}
	}

}

/* Function to sort map int -> int by its values */
func sortMapsByValueSize(m map[int]int) []int {
	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	return keys
}

func addDanglingNodes(m map[int][]int) {
	maxNode := -1
	for node := range m {
		if node > maxNode {
			maxNode = node
		}
	}
	for i := 0; i <= maxNode; i++ {
		if _, exists := m[i]; !exists {
			m[i] = []int{}
		}
	}
}

func randomSurf(graph map[int][]int, m float64, totalSteps int) {
	/* Initialize the map the keep track of nodes visited during the random surf. */
	randomSurferMap := make(map[int]int)
	for node := range graph {
		randomSurferMap[node] = 0
	}

	/*Indices of the nodes, to pick a random starting node - to avoid picking an index which doesnt exists. */
	nodes := make([]int, 0, len(graph))
	for node := range graph {
		nodes = append(nodes, node)
	}

	rand.Seed(time.Now().UnixNano())
	currentNode := nodes[rand.Intn(len(nodes))]

	for i := 0; i < totalSteps; i++ {
		randomSurferMap[currentNode]++

		/* If 1 > m < 0.15 or if node is dangling go to a random node, else follow a random adjacent edge of the current node. */
		if rand.Float64() < m || len(graph[currentNode]) == 0 {
			currentNode = nodes[rand.Intn(len(nodes))]
		} else {
			adjNodes := graph[currentNode]
			currentNode = adjNodes[rand.Intn(len(adjNodes))]
		}
	}

	sortedKeys := sortMapsByValueSize(randomSurferMap)

	fmt.Printf("====== Random surfer visits Top 10 ======\n")
	fmt.Printf("m: %f, steps: %d \n\n", m, totalSteps)
	fmt.Printf("Rank:		Node:		Count:\n")
	for i, key := range sortedKeys {
		if i >= 10 {
			break
		}
		fmt.Printf("%d		%d		%d\n", i+1, key, randomSurferMap[key])
	}
	fmt.Println()
}

func pageRank(graph map[int][]int, order int, size int) {

	/*Setup. Find branching factors, danglingNodes and reverse the graph. */
	branching := make(map[int]int)
	danglingNodes := []int{}
	for node, adjNodes := range graph {
		if len(adjNodes) == 0 {
			danglingNodes = append(danglingNodes, node)
		} else {
			branching[node] = len(adjNodes)
		}
	}
	sort.Slice(danglingNodes, func(i, j int) bool {
		return danglingNodes[i] > danglingNodes[j]
	})

	//fmt.Printf("dangling nodes (sorted): %d", danglingNodes)

	reversedGraph := make(map[int][]int)
	for node, adjNodes := range graph {
		for _, adjNode := range adjNodes {
			reversedGraph[adjNode] = append(reversedGraph[adjNode], node)
		}
	}

}
