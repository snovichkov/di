package di

import (
	"fmt"
)

// AliasMap is a slice of aliases.
type AliasMap map[string]string

// Copy returns a copy of the AliasMap.
func (a AliasMap) Copy() AliasMap {
	aliases := AliasMap{}

	for name, alias := range a {
		aliases[name] = alias
	}

	return aliases
}

// Get returns target if the AliasMap contains the given alias
// otherwise return original name.
func (a AliasMap) Get(name string) (string, error) {
	var checked = map[string]struct{}{}
	for {
		if _, ok := checked[name]; ok {
			return "", fmt.Errorf(`"%s" has circular reference of itself`, name)
		}

		if target, ok := a[name]; ok {
			name = target
			continue
		}

		break
	}

	return name, nil
}
