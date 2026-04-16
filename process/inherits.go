package process

import (
	"encoding/json"
	"fmt"
	"log"
)

func withInherits(p *Process) error {
	// Record child-level keys before any inheritance
	childKeys := make(map[string]struct{})
	tt := make(map[string]any)
	q, _ := json.Marshal(p)
	_ = json.Unmarshal(q, &tt)
	for k := range tt {
		childKeys[k] = struct{}{}
	}

	if len(p.Inherits) == 0 {
		return nil
	}

	tree := make([]map[string]any, 0)
	tree = append(tree, tt)

	next := p.Inherits
	provenance := make(map[string]string) // key -> source profile name

	for {
		parent, err := getSystemProcessRaw(next)
		if err != nil {
			return fmt.Errorf("error reading system process %s: %w", p.Inherits, err)
		}

		parentName := ""
		if n, ok := parent["name"].(string); ok {
			parentName = n
		}

		// Track provenance for keys not already in child
		for k := range parent {
			if _, seen := childKeys[k]; !seen {
				if _, recorded := provenance[k]; !recorded {
					provenance[k] = parentName
				}
			}
		}

		tree = append(tree, parent)
		v, ok := parent["inherits"].(string)
		if v == "" || !ok {
			break
		}
		next = v
	}

	newMap := make(map[string]any)
	for i := len(tree) - 1; i >= 0; i-- {
		for k, v := range tree[i] {
			newMap[k] = v
		}
	}

	jsonStr, _ := json.Marshal(newMap)
	err := json.Unmarshal(jsonStr, p)
	if err != nil {
		log.Println(err)
	}

	// Filter provenance to only include keys still present in final result
	// and exclude internal/metadata keys
	excludeKeys := map[string]struct{}{
		"name":              {},
		"from":              {},
		"inherits":          {},
		"version":           {},
		"is_custom_defined": {},
		"setting_id":        {},
		"print_settings_id": {},
	}

	p.InheritedFrom = make(map[string]string)
	for k, source := range provenance {
		if _, excluded := excludeKeys[k]; excluded {
			continue
		}
		if _, present := newMap[k]; present {
			p.InheritedFrom[k] = source
		}
	}

	return nil
}
