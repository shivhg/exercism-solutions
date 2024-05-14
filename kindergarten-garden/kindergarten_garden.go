package kindergarten

import (
	"errors"
	"slices"
	"strings"
)

// Define the Garden type here.

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

type Garden struct {
	Diagram  string
	Children []string
	CupCodes map[string]string
}

func NewGarden(diagram string, ch []string) (*Garden, error) {
	children := slices.Clone(ch)
	slices.Sort(children)

	ss := strings.Split(diagram, "\n")
	if len(ss) != 3 || len(ss[1]) != len(ss[2]) || len(ss[1])%2 != 0 {
		return nil, errors.New("Err")
	}
	chilLen := len(children)
	children = slices.Compact(children)
	if chilLen != len(children) {
		return nil, errors.New("Err")
	}

	gar := Garden{Diagram: diagram, Children: children, CupCodes: map[string]string{
		"C": "clover",
		"G": "grass",
		"R": "radishes",
		"V": "violets",
	}}

	for i, _ := range ss[1] {
		_, exist := gar.CupCodes[ss[1][i:i+1]]
		if !exist {
			return nil, errors.New("Err")
		}
	}
	return &gar, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	if !slices.Contains(g.Children, child) {
		return []string{}, false
	}

	childInd := slices.Index(g.Children, child)
	childLookupIndex := childInd * 2

	ss := strings.Split(g.Diagram, "\n")
	row1 := ss[1]
	row2 := ss[2]

	row1Sel := row1[childLookupIndex : childLookupIndex+2]
	row2Sel := row2[childLookupIndex : childLookupIndex+2]

	return []string{g.CupCodes[row1Sel[0:1]], g.CupCodes[row1Sel[1:2]], g.CupCodes[row2Sel[0:1]], g.CupCodes[row2Sel[1:2]]}, true
}
