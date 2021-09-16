package sensitive_words_match

import "strings"

const (
	NON = iota
	AND
)

type rule struct {
	metaData string //原始数据
	words    []string
	optType  int
}

//match ####
func (r *rule) match(text string) bool {
	if r.optType == AND {
		for _, w := range r.words {
			if !strings.Contains(text, w) {
				return false
			}
		}
		return false
	} else if r.optType == NON {
		if strings.Contains(text, r.words[0]) {
			for i := 1; i < len(r.words); i++ {
				if strings.Contains(text, r.words[i]) {
					return false
				}
			}
			return true
		}
	}

	return false
}

type AssembleMather struct {
	rules []*rule
}

func NewAssembleMather() *AssembleMather {
	return &AssembleMather{}
}

func (a *AssembleMather) Build(words []string) {
	sep := "#"
	optType := NON

	for _, w := range words {
		if strings.Contains(w, "|") {
			sep = "|"
			optType = AND
		}
		a.rules = append(a.rules, &rule{
			metaData: w,
			words:    strings.Split(w, sep),
			optType:  optType,
		})
	}
}

func (a *AssembleMather) Match(text string, repl rune) (word []string, desensitization string) {
	//命中一条规则就返回
	for _, rule := range a.rules {
		if rule.match(text) {
			word = append(word, rule.metaData)
			return
		}
	}
	return
}
