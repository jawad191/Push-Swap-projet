package piles

import (
	"errors"
	"strconv"
	"strings"
)

type Piles struct {
	A          []int
	B          []int
	Operations []string
}

// ParseArgs converts arguments into integers and checks duplicates.
func ParseArgs(args []string) ([]int, error) {
	if len(args) == 0 {
		return nil, nil
	}

	var tokens []string

	for _, arg := range args {
		fields := strings.Fields(arg)

		if len(fields) == 0 {
			return nil, errors.New("empty argument")
		}

		tokens = append(tokens, fields...)
	}

	seen := make(map[int]bool)
	nums := make([]int, 0, len(tokens))

	for _, tok := range tokens {
		val, err := strconv.Atoi(tok)
		if err != nil {
			return nil, errors.New("invalid integer")
		}

		if seen[val] {
			return nil, errors.New("duplicate integer")
		}

		seen[val] = true
		nums = append(nums, val)
	}

	return nums, nil
}

// IsSorted checks whether A is sorted in ascending order.
func (p *Piles) IsSorted() bool {
	for i := 0; i+1 < len(p.A); i++ {
		if p.A[i] > p.A[i+1] {
			return false
		}
	}
	return true
}

// Sa swaps the first two elements of A.
func (p *Piles) Sa() {
	if len(p.A) < 2 {
		return
	}

	p.A[0], p.A[1] = p.A[1], p.A[0]
	p.Operations = append(p.Operations, "sa")
}

// Sb swaps the first two elements of B.
func (p *Piles) Sb() {
	if len(p.B) < 2 {
		return
	}

	p.B[0], p.B[1] = p.B[1], p.B[0]
	p.Operations = append(p.Operations, "sb")
}

// Ss swaps the first two elements of both stacks.
func (p *Piles) Ss() {
	if len(p.A) < 2 && len(p.B) < 2 {
		return
	}

	if len(p.A) >= 2 {
		p.A[0], p.A[1] = p.A[1], p.A[0]
	}

	if len(p.B) >= 2 {
		p.B[0], p.B[1] = p.B[1], p.B[0]
	}

	p.Operations = append(p.Operations, "ss")
}

// Pa pushes the first element of B onto A.
func (p *Piles) Pa() {
	if len(p.B) == 0 {
		return
	}

	p.A = append([]int{p.B[0]}, p.A...)
	p.B = p.B[1:]

	p.Operations = append(p.Operations, "pa")
}

// Pb pushes the first element of A onto B.
func (p *Piles) Pb() {
	if len(p.A) == 0 {
		return
	}

	p.B = append([]int{p.A[0]}, p.B...)
	p.A = p.A[1:]

	p.Operations = append(p.Operations, "pb")
}

// Ra rotates A upwards.
func (p *Piles) Ra() {
	if len(p.A) < 2 {
		return
	}

	p.A = append(p.A[1:], p.A[0])
	p.Operations = append(p.Operations, "ra")
}

// Rb rotates B upwards.
func (p *Piles) Rb() {
	if len(p.B) < 2 {
		return
	}

	p.B = append(p.B[1:], p.B[0])
	p.Operations = append(p.Operations, "rb")
}

// Rr rotates both stacks upwards.
func (p *Piles) Rr() {
	if len(p.A) < 2 && len(p.B) < 2 {
		return
	}

	if len(p.A) >= 2 {
		p.A = append(p.A[1:], p.A[0])
	}

	if len(p.B) >= 2 {
		p.B = append(p.B[1:], p.B[0])
	}

	p.Operations = append(p.Operations, "rr")
}

// Rra rotates A downwards.
func (p *Piles) Rra() {
	if len(p.A) < 2 {
		return
	}

	last := p.A[len(p.A)-1]
	p.A = append([]int{last}, p.A[:len(p.A)-1]...)

	p.Operations = append(p.Operations, "rra")
}

// Rrb rotates B downwards.
func (p *Piles) Rrb() {
	if len(p.B) < 2 {
		return
	}

	last := p.B[len(p.B)-1]
	p.B = append([]int{last}, p.B[:len(p.B)-1]...)

	p.Operations = append(p.Operations, "rrb")
}

// Rrr rotates both stacks downwards.
func (p *Piles) Rrr() {
	if len(p.A) < 2 && len(p.B) < 2 {
		return
	}

	if len(p.A) >= 2 {
		last := p.A[len(p.A)-1]
		p.A = append([]int{last}, p.A[:len(p.A)-1]...)
	}

	if len(p.B) >= 2 {
		last := p.B[len(p.B)-1]
		p.B = append([]int{last}, p.B[:len(p.B)-1]...)
	}

	p.Operations = append(p.Operations, "rrr")
}

// Exec executes an instruction.
func (p *Piles) Exec(op string) error {
	switch op {
	case "pa":
		p.Pa()
	case "pb":
		p.Pb()
	case "sa":
		p.Sa()
	case "sb":
		p.Sb()
	case "ss":
		p.Ss()
	case "ra":
		p.Ra()
	case "rb":
		p.Rb()
	case "rr":
		p.Rr()
	case "rra":
		p.Rra()
	case "rrb":
		p.Rrb()
	case "rrr":
		p.Rrr()
	default:
		return errors.New("invalid instruction")
	}

	return nil
}
