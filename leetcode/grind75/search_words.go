package grind75

// 211. Design Add and Search Words Data Structure

// Design a data structure that supports adding new words and finding if a string matches any previously added string.

// Implement the WordDictionary class:

// WordDictionary() Initializes the object.
// void addWord(word) Adds word to the data structure, it can be matched later.
// bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.

// Example:

// Input
// ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
// [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
// Output
// [null,null,null,null,false,true,true,true]

// Explanation
// WordDictionary wordDictionary = new WordDictionary();
// wordDictionary.addWord("bad");
// wordDictionary.addWord("dad");
// wordDictionary.addWord("mad");
// wordDictionary.search("pad"); // return False
// wordDictionary.search("bad"); // return True
// wordDictionary.search(".ad"); // return True
// wordDictionary.search("b.."); // return True

// Constraints:

// 1 <= word.length <= 25
// word in addWord consists of lowercase English letters.
// word in search consist of '.' or lowercase English letters.
// There will be at most 3 dots in word for search queries.
// At most 104 calls will be made to addWord and search.

import (
	"container/list"
)

type WordTrieNode struct {
	IsTerminal bool
	Keys       map[rune]*WordTrieNode
	Prefix     string
}

type WordDictionary struct {
	root *WordTrieNode
}

func NewWordDictionary() WordDictionary {
	return WordDictionary{
		root: &WordTrieNode{
			IsTerminal: false,
			Prefix:     "",
			Keys:       map[rune]*WordTrieNode{},
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
	currNode := this.root
	for i, chr := range word {
		if _, found := currNode.Keys[chr]; !found {
			currNode.Keys[chr] = &WordTrieNode{
				Prefix: word[:i+1],
				Keys:   map[rune]*WordTrieNode{},
			}
		}
		currNode = currNode.Keys[chr]
	}
	currNode.IsTerminal = true
}

func (this *WordDictionary) Search(word string) bool {
	queue := list.New()
	queue.PushBack(this.root)
	var currWordIndex int
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			e := queue.Front()
			node := e.Value.(*WordTrieNode)
			if currWordIndex == len(word) {
				if node.IsTerminal {
					return true
				}
			} else {
				letter := rune(word[currWordIndex])
				if letter == '.' {
					for _, node := range node.Keys {
						queue.PushBack(node)
					}
				} else {
					if _, found := node.Keys[letter]; found {
						queue.PushBack(node.Keys[letter])
					}
				}
			}
			queue.Remove(e)
		}
		currWordIndex++
	}
	return false
}
