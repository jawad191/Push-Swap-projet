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

// sa swaps the first two elements of A.
func (p *Piles) sa() {
	if len(p.A) < 2 {
		return
	}

	p.A[0], p.A[1] = p.A[1], p.A[0]
	p.Operations = append(p.Operations, "sa")
}

// sb swaps the first two elements of B.
func (p *Piles) sb() {
	if len(p.B) < 2 {
		return
	}

	p.B[0], p.B[1] = p.B[1], p.B[0]
	p.Operations = append(p.Operations, "sb")
}

// ss swaps the first two elements of both stacks.
func (p *Piles) ss() {
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

// pa pushes the first element of B onto A.
func (p *Piles) pa() {
	if len(p.B) == 0 {
		return
	}

	p.A = append([]int{p.B[0]}, p.A...)
	p.B = p.B[1:]

	p.Operations = append(p.Operations, "pa")
}

// pb pushes the first element of A onto B.
func (p *Piles) pb() {
	if len(p.A) == 0 {
		return
	}

	p.B = append([]int{p.A[0]}, p.B...)
	p.A = p.A[1:]

	p.Operations = append(p.Operations, "pb")
}

// ra rotates A upwards.
func (p *Piles) ra() {
	if len(p.A) < 2 {
		return
	}

	p.A = append(p.A[1:], p.A[0])
	p.Operations = append(p.Operations, "ra")
}

// rb rotates B upwards.
func (p *Piles) rb() {
	if len(p.B) < 2 {
		return
	}

	p.B = append(p.B[1:], p.B[0])
	p.Operations = append(p.Operations, "rb")
}

// rr rotates both stacks upwards.
func (p *Piles) rr() {
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

// rra rotates A downwards.
func (p *Piles) rra() {
	if len(p.A) < 2 {
		return
	}

	last := p.A[len(p.A)-1]
	p.A = append([]int{last}, p.A[:len(p.A)-1]...)

	p.Operations = append(p.Operations, "rra")
}

// rrb rotates B downwards.
func (p *Piles) rrb() {
	if len(p.B) < 2 {
		return
	}

	last := p.B[len(p.B)-1]
	p.B = append([]int{last}, p.B[:len(p.B)-1]...)

	p.Operations = append(p.Operations, "rrb")
}

// rrr rotates both stacks downwards.
func (p *Piles) rrr() {
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
