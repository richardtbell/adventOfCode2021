package day14

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func (p *Polymer) addToPolymerChain() {
	newChain := ""
	for i := 0; i < len(p.chain)-1; i++ {
		pair := string(p.chain[i : i+2])
		newChain += string(p.chain[i]) + p.insertion[pair]
	}
	p.chain = newChain + string(p.chain[len(p.chain)-1])
}

func (p *Polymer) addNewPairs() {
	newPairs := map[string]int{}
	for pair, count := range p.pairs {
		mid := p.insertion[pair]
		newPair1 := string(pair[0]) + mid
		newPair2 := mid + string(pair[1])
		newPairs[newPair1] += count
		newPairs[newPair2] += count
	}
	p.pairs = newPairs
}

type Polymer struct {
	insertion map[string]string
	chain     string
	pairs     map[string]int
}

func readFile(filepath string) Polymer {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	both := strings.Split(string(contents), "\n\n")
	p := Polymer{chain: both[0], insertion: map[string]string{}, pairs: map[string]int{}}
	for i := 0; i < len(p.chain)-1; i++ {
		pair := string(p.chain[i : i+2])
		p.pairs[pair]++
	}
	insertionRules := strings.Split(string(both[1]), "\n")
	for _, rule := range insertionRules {
		r := strings.Split(string(rule), " -> ")
		p.insertion[r[0]] = r[1]
	}
	return p
}

func (p Polymer) getMaxAndMin() (int, int) {
	m := make(map[string]int, 0)
	for _, c := range p.chain {
		m[string(c)]++
	}
	fmt.Println(m)
	vals := make([]int, 0, len(m))
	for _, v := range m {
		vals = append(vals, v)
	}
	sort.Ints(vals)

	return vals[len(vals)-1], vals[0]
}

func CalculatePart1(path string) int {
	p := readFile(path)
	for i := 0; i < 10; i++ {
		p.addToPolymerChain()
	}
	max, min := p.getMaxAndMin()
	return max - min
}

func (p Polymer) getMaxAndMinFromPairs() (int, int) {
	vals := make([]int, 0, len(p.pairs))
	letters := map[string]int{}
	for k, v := range p.pairs {
		letters[string(k[0])] += v
	}
	letters[string(p.chain[len(p.chain)-1])]++

	for _, v := range letters {
		vals = append(vals, v)
	}
	sort.Ints(vals)

	return vals[len(vals)-1], vals[0]
}

func CalculatePart2(path string) int {
	p := readFile(path)
	for i := 0; i < 40; i++ {
		p.addNewPairs()
	}

	max, min := p.getMaxAndMinFromPairs()
	return max - min
}
