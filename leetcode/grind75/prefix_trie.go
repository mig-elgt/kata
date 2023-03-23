package grind75

// 208. Implement Trie (Prefix Tree)

// A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

// Implement the Trie class:

// Trie() Initializes the trie object.
// void insert(String word) Inserts the string word into the trie.
// boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
// boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.

// Example 1:
// Input
// ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
// [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
// Output
// [null, null, true, false, true, null, true]

// Explanation
// Trie trie = new Trie();
// trie.insert("apple");
// trie.search("apple");   // return True
// trie.search("app");     // return False
// trie.startsWith("app"); // return True
// trie.insert("app");
// trie.search("app");     // return True

// Constraints:

// 1 <= word.length, prefix.length <= 2000
// word and prefix consist only of lowercase English letters.
// At most 3 * 104 calls in total will be made to insert, search, and startsWith.

type TrieNode struct {
	prefix     string
	keys       map[rune]*TrieNode
	isTerminal bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() Trie {
	return Trie{root: &TrieNode{
		prefix: "",
		keys:   map[rune]*TrieNode{},
	}}
}

func (this *Trie) Insert(word string) {
	currNode := this.root
	for i := 0; i < len(word); i++ {
		key := rune(word[i])
		if _, found := currNode.keys[key]; !found {
			child := &TrieNode{
				prefix: string(word[:i+1]),
				keys:   map[rune]*TrieNode{},
			}
			currNode.keys[key] = child
		}
		currNode = currNode.keys[key]
	}
	currNode.isTerminal = true
}

func (this *Trie) Search(word string) bool {
	return this.recSearch(this.root, word)
}

func (this Trie) recSearch(node *TrieNode, word string) bool {
	if word == "" {
		if node.isTerminal {
			return true
		}
		return false
	}
	key := rune(word[0])
	if _, ok := node.keys[key]; !ok {
		return false
	}
	return this.recSearch(node.keys[key], word[1:])
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for _, p := range prefix {
		if _, found := curr.keys[p]; !found {
			return false
		}
		curr = curr.keys[p]
	}
	return true
}
