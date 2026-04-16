package process

import (
	"encoding/json"
	"fmt"
	"log"
)

func withInherits(p *Process) error {
	tree := make([]map[string]any, 0)

	fmt.Printf("%s\n", p.Name)

	tt := make(map[string]any)

	// TODO error
	q, _ := json.Marshal(p)
	_ = json.Unmarshal(q, &tt)

	tree = append(tree, tt)

	if len(p.Inherits) == 0 {
		return nil
	}

	next := p.Inherits

	for {
		parent, err := getSystemProcessRaw(next)
		if err != nil {
			return fmt.Errorf("error reading system process %s: %w", p.Inherits, err)
		}

		tree = append(tree, parent)
		v, ok := parent["inherits"].(string)
		if v == "" || !ok {
			break
		}
		next = v
	}

	// for _, e := range tree {
	// 	fmt.Printf(" -> %s", e["name"])
	// }
	// fmt.Println()

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
		// return nil, err
	}

	return nil
}
