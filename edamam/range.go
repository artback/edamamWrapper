package edamam

import "fmt"

type Range struct {
	To   int
	From int
}

func (r Range) String() string {
	if r.From == 0 && r.To == 0 {
		return ""
	}
	if r.To == 0 {
		return fmt.Sprintf("%d+", r.From)
	} else if r.From == 0 {
		return fmt.Sprintf("%d", r.To)
	}
	return fmt.Sprintf("%d-%d", r.From, r.To)
}
