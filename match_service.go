package sensitive_words_match

import "strings"

const (
	DFA      = 0
	ASSEMBLE = 1
	REGEXP   = 2
)

type MatchService struct {
	matchers map[int]Matcher
}

func NewMatchService() *MatchService {
	return &MatchService{
		matchers: make(map[int]Matcher),
	}
}

//Build 同一个配置中支持三种配置
func (m *MatchService) Build(words []string) {
	var (
		dfaList      []string
		assembleList []string
		regexpList   []string
	)

	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "reg@") {
			regexpList = append(regexpList, words[i][len("reg@"):])
		} else if strings.Contains(words[i], "|") || strings.Contains(words[i], "#") {
			assembleList = append(assembleList, words[i])
		} else {
			dfaList = append(dfaList, words[i])
		}
	}

	if len(dfaList) > 0 {
		matcher := NewDFAMather()
		matcher.Build(dfaList)
		m.matchers[DFA] = matcher
	}

	if len(assembleList) > 0 {
		matcher := NewAssembleMather()
		matcher.Build(assembleList)
		m.matchers[ASSEMBLE] = matcher
	}

	if len(regexpList) > 0 {
		matcher := NewRegexpMatcher()
		matcher.Build(regexpList)
		m.matchers[REGEXP] = matcher
	}
}

//Match 只要有一个规则就返回；
func (m *MatchService) Match(text string, repl rune) (sensitiveWords []string, replaceText string) {
	for _, x := range m.matchers {
		sensitiveWords, replaceText = x.Match(text, '*')
		if len(sensitiveWords) > 0 {
			return
		}
	}
	return
}
