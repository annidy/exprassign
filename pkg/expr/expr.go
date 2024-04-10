package expr

import (
	"regexp"
	"strings"
)

type AssignExpr struct {
	Key   string
	Value string
}

func (e AssignExpr) Assign(s string) (r string, ok bool) {
	var re = regexp.MustCompile(`(\s|^)` + e.Key + `(\s*)(:)?=(\s*)(\b(\d+)\b|\b(\d+\.\d+)\b|"(?:[^"\\]|\\.)*")`)
	for _, matche := range re.FindAllStringSubmatch(s, -1) {
		r = strings.Replace(s, matche[5], e.Value, 1)
		ok = true
	}
	return
}
