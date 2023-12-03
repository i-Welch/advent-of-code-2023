package trie

type Trie struct {
	children     map[rune]Trie
	word         bool
	asciiNumeral rune
}

func (t Trie) Add(word string, numeral rune) {
	root := t
	for _, char := range word {
		child, ok := root.children[char]
		if !ok {
			child = Trie{}
			root.children[char] = child
		}
		root = child
	}
	root.word = true
	root.asciiNumeral = numeral
}

func (t Trie) Get(word string) (rune, bool) {
	root := t
	for _, char := range word {
		child, ok := root.children[char]
		if !ok {
			return '\x00', false
		}
		root = child
	}

	return root.asciiNumeral, root.word
}
