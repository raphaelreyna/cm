package maps

import (
	"errors"
	"sort"
)

var (
	maps map[string]*Map

	ErrInvalidCharSys  = errors.New("invalid character system")
	ErrInvalidCharName = errors.New("invalid character name")
)

type Char struct {
	U rune
	L rune
}

type Map struct {
	Name string

	Chars map[string]*Char

	sortedNames []string
}

func (m *Map) SortedNames() []string {
	if m.sortedNames != nil {
		return m.sortedNames
	}

	idx := 0
	m.sortedNames = make([]string, len(m.Chars))
	for name := range m.Chars {
		m.sortedNames[idx] = name
		idx++
	}

	sort.Strings(m.sortedNames)

	return m.sortedNames
}

func MapChar(charSys, charName string) (*Char, error) {
	m, ok := maps[charSys]
	if !ok {
		return nil, ErrInvalidCharSys
	}

	r, ok := m.Chars[charName]
	if !ok {
		return nil, ErrInvalidCharName
	}

	return r, nil
}

func GetMap(charSys string) (*Map, error) {
	m, ok := maps[charSys]
	if !ok {
		return nil, ErrInvalidCharSys
	}

	return m, nil
}

func All() map[string]*Map {
	return maps
}

func Count() int {
	return len(maps)
}
