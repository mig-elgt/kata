package grind75

import (
	"container/list"
	"strconv"
)

func calculate(s string) int {
	tokens := list.New()
	values := list.New()
	for i := 0; i < len(s); i++ {
		chr := s[i]
		if chr != ' ' {
			var number string
			j := i
			for ; j < len(s); j++ {
				if s[j] >= 48 && s[j] <= 57 {
					number += string(s[j])
				} else {
					break
				}
			}
			if len(number) > 0 {
				num, _ := strconv.Atoi(number)
				values.PushBack(num)
				i = j - 1
				continue
			}
			if chr == '(' {
				tokens.PushBack(rune(chr))
			} else if chr == '+' || chr == '-' {
				for {
					element := tokens.Back()
					if element != nil {
						if element.Value.(rune) == '-' {
							tokens.Remove(element)
							eval(element.Value.(rune), values)
						} else {
							tokens.PushBack(rune(chr))
							break
						}
					} else {
						tokens.PushBack(rune(chr))
						break
					}
				}
			} else if chr == ')' {
				for {
					element := tokens.Back()
					tokens.Remove(element)
					if element.Value.(rune) != '(' {
						eval(element.Value.(rune), values)
					} else {
						break
					}
				}
			}
		}
	}
	for tokens.Len() > 0 {
		element := tokens.Back()
		tokens.Remove(element)
		eval(element.Value.(rune), values)
	}
	return values.Back().Value.(int)
}

func eval(operator rune, stackValues *list.List) {
	b := stackValues.Back()
	stackValues.Remove(b)
	a := stackValues.Back()
	stackValues.Remove(a)
	res := operators[operator](a.Value.(int), b.Value.(int))
	stackValues.PushBack(res)
}

var operators = map[rune]func(a, b int) int{
	'+': func(a, b int) int {
		return a + b
	},
	'-': func(a, b int) int {
		return a - b
	},
}
