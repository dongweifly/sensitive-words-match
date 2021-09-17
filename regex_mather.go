package sensitive_words_match

import (
	"fmt"
	"github.com/dlclark/regexp2"
)

type regRule struct {
	metaData string
	reg      *regexp2.Regexp
}

func NewRegRule(str string) (*regRule, error) {
	if r, err := regexp2.Compile(str, 0); err == nil {
		return &regRule{
			metaData: str,
			reg:      r,
		}, nil
	} else {
		return nil, fmt.Errorf("%s comiple regexp error:%s ", str, err.Error())
	}
}

func (r *regRule) MatchAll(text string) []string {
	var ret []string
	match, _ := r.reg.FindStringMatch(text)
	if match != nil {
		ret = append(ret, match.String())
		match, _ = r.reg.FindNextMatch(match)
	}
	return ret
}

type RegexpMather struct {
	matchers []*regRule
}

func NewRegexpMatcher() *RegexpMather {
	return &RegexpMather{
	}
}

func (a *RegexpMather) Build(words []string) {
	for _, w := range words {
		if m, err := NewRegRule(w); err == nil {
			a.matchers = append(a.matchers, m)
		}
	}
	return
}

//Match 所有的正则只有命中一个就返回；
func (a *RegexpMather) Match(text string, repl rune) (word []string, desensitization string) {
	desensitization = text
	for _, r := range a.matchers {
		word = r.MatchAll(text)
		if len(word) > 0 {
			return
		}
	}
	return
}
