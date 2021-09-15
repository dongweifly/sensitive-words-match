package sensitive_words_match

type Matcher interface {
	//Build build Matcher
	Build(words []string)

	//Match return match sensitive words
	Match(text string, repl rune) ([]string, string)
}
