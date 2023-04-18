package bingtranslate

import (
	"fmt"
)

type translationWord struct {
	target      string
	backTargets []string
}

func (t *translationWord) add(s string) {
	t.backTargets = append(t.backTargets, s)
}

type posSet map[string][]translationWord

func (set posSet) add(tag string, words translationWord) {
	if _, ok := set[tag]; !ok {
		set[tag] = []translationWord{words}
	} else {
		set[tag] = append(set[tag], words)
	}
}

func (set posSet) format() (s string) {
	for tag := range set {
		s += fmt.Sprintf("[%s]\n", tag)
		for _, words := range set[tag] {
			s += fmt.Sprintf("\t%s:", words.target)
			firstWord := true
			for _, backTarget := range words.backTargets {
				if firstWord {
					s += fmt.Sprintf(" %s", backTarget)
					firstWord = false
				} else {
					s += fmt.Sprintf(", %s", backTarget)
				}
			}
			s += "\n"
		}
	}
	return s
}
