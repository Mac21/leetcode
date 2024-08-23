package main

import (
	"strings"
)

/*
10. Regular Expression Matching Hard
Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Example 1:

Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".


Constraints:

1 <= s.length <= 20
1 <= p.length <= 20
s contains only lowercase English letters.
p contains only lowercase English letters, '.', and '*'.
It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
*/

type State struct {
	Label byte
	Left  *State
	Right *State
}

func follows_helper(s *State, cache *map[*State]bool) {
	if s == nil {
		return
	}

	if (*cache)[s] {
		return
	}

	(*cache)[s] = true

	if s.Label == 0 {
		if s.Left != nil {
			follows_helper(s.Left, cache)
		}
		if s.Right != nil {
			follows_helper(s.Right, cache)
		}
	}
}

type Regex struct {
	Initial *State
	Accept  *State
}

func (r Regex) Match(in string) bool {
	current := make(map[*State]bool)
	next := make(map[*State]bool)
	follows_helper(r.Initial, &current)
	for i, c := range in {
		for s := range current {
			if s.Label == '.' || s.Label == byte(c) {
				follows_helper(s.Left, &next)
			}
		}
		if i < len(in)-1 {
			temp := current
			current = next
            clear(temp)
			next = temp
		}
	}

	return next[r.Accept]
}

// Transforms re from ab to a+b so that shunting yard algorithm has concat op
// also apply a minimal optimization by removing any consecutive klenee star
// e.g. a***b -> a*+b
func transform(re string) string {
	var sb strings.Builder
	sb.WriteByte(re[0])
	for i := 1; i < len(re); i++ {
		c := re[i]
		switch c {
		case '*':
			// If we're seeing a * and the previous character was a * skip.
			if re[i-1] == '*' {
				continue
			}

			sb.WriteByte(c)
		default:
			sb.WriteByte('+')
			sb.WriteByte(c)
		}
	}
	return sb.String()
}

func shunt(re string) string {
	precedence := map[rune]int{
		'*': 60,
		// concatenation
		'+': 40,
	}
	re = transform(re)
	var output strings.Builder
	ops := make([]rune, 0, len(re))
	for _, c := range re {
		switch {
		case precedence[c] > 0:
			for len(ops) > 0 && precedence[ops[len(ops)-1]] > precedence[c] {
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]
				output.WriteRune(op)
			}
			ops = append(ops, c)
		default:
			output.WriteRune(c)
		}
	}
	for len(ops) > 0 {
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		output.WriteRune(op)
	}
	return output.String()
}

func compile(postfix string) *Regex {
	nfas := make([]*Regex, 0)
	for _, c := range postfix {
		switch c {
		case '*':
			re := nfas[len(nfas)-1]
			nfas[len(nfas)-1] = nil
			nfas = nfas[:len(nfas)-1]

			a := new(State)
			i := &State{
				Left:  re.Initial,
				Right: a,
			}
			re.Accept.Left = re.Initial
			re.Accept.Right = a

			nfas = append(nfas, &Regex{
				Initial: i,
				Accept:  a,
			})
		case '+':
			re2 := nfas[len(nfas)-1]
			nfas[len(nfas)-1] = nil
			nfas = nfas[:len(nfas)-1]

			re1 := nfas[len(nfas)-1]
			nfas[len(nfas)-1] = nil
			nfas = nfas[:len(nfas)-1]

			re1.Accept.Left = re2.Initial

			nfas = append(nfas, &Regex{
				Initial: re1.Initial,
				Accept:  re2.Accept,
			})
		default:
			a := new(State)
			i := &State{
				Label: byte(c),
				Left:  a,
			}
			nfas = append(nfas, &Regex{
				Initial: i,
				Accept:  a,
			})
		}
	}
	return nfas[len(nfas)-1]
}

func NewRegex(re string) *Regex {
	postfix := shunt(re)
	return compile(postfix)
}

func isMatch(s, p string) bool {
	re := NewRegex(p)
	return re.Match(s)
}