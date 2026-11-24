package piles

func (p *Piles) Sort() {
	if len(p.A) <= 1 || p.IsSorted() {
		return
	}
	if len(p.A) == 2 {
		p.Sa()
		return
	}
	if len(p.A) == 3 {
		p.sortThree()
		return
	}

	// Push 2 elements to B to establish min/max bounds in B
	if len(p.A) > 3 && !p.IsSorted() {
		p.Pb()
	}
	if len(p.A) > 3 && !p.IsSorted() {
		p.Pb()
	}

	// Push elements from A to B in descending order using cheapest combined rotation
	for len(p.A) > 3 && !p.IsSorted() {
		p.moveCheapestAToB()
	}

	p.sortThree()

	// Push all elements back from B into their sorted positions in A
	for len(p.B) > 0 {
		p.moveBToA()
	}

	// Rotate minimum element of A to the top
	p.minToTop()
}

func (p *Piles) sortThree() {
	x, y, z := p.A[0], p.A[1], p.A[2]
	switch {
	case x <= y && y <= z:
	case x <= z && z < y:
		p.Sa()
		p.Ra()
	case y < x && x <= z:
		p.Sa()
	case z < x && x <= y:
		p.Rra()
	case y < z && z < x:
		p.Ra()
	default:
		p.Sa()
		p.Rra()
	}
}

func (p *Piles) minIndex(s []int) int {
	idx := 0
	for i := 1; i < len(s); i++ {
		if s[i] < s[idx] {
			idx = i
		}
	}
	return idx
}

func (p *Piles) maxIndex(s []int) int {
	idx := 0
	for i := 1; i < len(s); i++ {
		if s[i] > s[idx] {
			idx = i
		}
	}
	return idx
}

// targetInB finds the index in descending stack B where val should sit at the top before pb.
func (p *Piles) targetInB(val int) int {
	bestIdx := -1
	bestVal := 0
	for i, bVal := range p.B {
		if bVal < val && (bestIdx == -1 || bVal > bestVal) {
			bestIdx = i
			bestVal = bVal
		}
	}
	if bestIdx != -1 {
		return bestIdx
	}
	return p.maxIndex(p.B)
}

// targetInA finds the index in ascending stack A where val should sit at the top before pa.
func (p *Piles) targetInA(val int) int {
	bestIdx := -1
	bestVal := 0
	for i, aVal := range p.A {
		if aVal > val && (bestIdx == -1 || aVal < bestVal) {
			bestIdx = i
			bestVal = aVal
		}
	}
	if bestIdx != -1 {
		return bestIdx
	}
	return p.minIndex(p.A)
}

func calcCost(idxA, lenA, idxB, lenB int) (cost, rotA, rotB int) {
	// rotA > 0 means ra, rotA < 0 means rra
	if idxA <= lenA/2 {
		rotA = idxA
	} else {
		rotA = idxA - lenA
	}
	if idxB <= lenB/2 {
		rotB = idxB
	} else {
		rotB = idxB - lenB
	}

	if rotA >= 0 && rotB >= 0 {
		if rotA > rotB {
			return rotA, rotA, rotB
		}
		return rotB, rotA, rotB
	}
	if rotA < 0 && rotB < 0 {
		absA, absB := -rotA, -rotB
		if absA > absB {
			return absA, rotA, rotB
		}
		return absB, rotA, rotB
	}
	absA, absB := rotA, rotB
	if absA < 0 {
		absA = -absA
	}
	if absB < 0 {
		absB = -absB
	}
	return absA + absB, rotA, rotB
}

func (p *Piles) applyRotations(rotA, rotB int) {
	for rotA > 0 && rotB > 0 {
		p.Rr()
		rotA--
		rotB--
	}
	for rotA < 0 && rotB < 0 {
		p.Rrr()
		rotA++
		rotB++
	}
	for rotA > 0 {
		p.Ra()
		rotA--
	}
	for rotA < 0 {
		p.Rra()
		rotA++
	}
	for rotB > 0 {
		p.Rb()
		rotB--
	}
	for rotB < 0 {
		p.Rrb()
		rotB++
	}
}

func (p *Piles) moveCheapestAToB() {
	bestCost := int(^uint(0) >> 1)
	bestRotA, bestRotB := 0, 0

	for i, val := range p.A {
		tB := p.targetInB(val)
		cost, rA, rB := calcCost(i, len(p.A), tB, len(p.B))
		if cost < bestCost {
			bestCost = cost
			bestRotA = rA
			bestRotB = rB
		}
	}

	p.applyRotations(bestRotA, bestRotB)
	p.Pb()
}

func (p *Piles) moveBToA() {
	tA := p.targetInA(p.B[0])
	if tA <= len(p.A)/2 {
		for i := 0; i < tA; i++ {
			p.Ra()
		}
	} else {
		for i := 0; i < len(p.A)-tA; i++ {
			p.Rra()
		}
	}
	p.Pa()
}

func (p *Piles) minToTop() {
	minIdx := p.minIndex(p.A)
	if minIdx <= len(p.A)/2 {
		for i := 0; i < minIdx; i++ {
			p.Ra()
		}
	} else {
		for i := 0; i < len(p.A)-minIdx; i++ {
			p.Rra()
		}
	}
}