package main

import "fmt"
import "math/rand"
import "time"

const NUM_CELLS = 64

func main() {
	fmt.Println("Initialize population...")
	rand.Seed(time.Now().UTC().UnixNano())
	var automata uint64 = 0x00000001
	for i := 0; i < 250; i++ {
		fmt.Printf("%064b\n", automata)
		automata = generateNextAutomata(automata)
	}
}

func generateNextAutomata(prevAutomata uint64) uint64 {
	returnAutomata := prevAutomata
	for i := NUM_CELLS - 1; i >= 0; i-- {
		middleAutomata := (prevAutomata >> uint(i)) & 0x00000001
		leftAutomata := prevAutomata >> uint(i+1)
		rightAutomata := prevAutomata >> uint(i-1)
		middle := int(middleAutomata & 0x00000001)
		left := int(leftAutomata & 0x00000001)
		right := int(rightAutomata & 0x00000001)

		newState := calculateAutomataBitState(left, middle, right)
		if newState == 1 {
			returnAutomata |= (1 << uint(i))
		} else {
			returnAutomata &^= (1 << uint(i))
		}
	}
	return returnAutomata
}

func calculateAutomataBitState(left, middle, right int) int {
	autoCase := fmt.Sprintf("%d%d%d", left, middle, right)
	var bit int
	switch autoCase {
	case "000":
		bit = 0
	case "001":
		bit = 1
	case "010":
		bit = 1
	case "011":
		bit = 1
	case "100":
		bit = 0
	case "101":
		bit = 1
	case "110":
		bit = 1
	case "111":
		bit = 0
	}
	return bit
}
