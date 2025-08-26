package entity

import (
	"encoding/json"
	"fmt"
	"os"
)

type (
	JSONParser struct {
		Data map[string]any
	}
)

func (p *JSONParser) DeserializeJSON(path string) (map[string]any, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	if err := json.Unmarshal(raw, &p.Data); err != nil {
		return nil, fmt.Errorf("unmarshal json: %w", err)
	}

	return p.Data, nil
}

func (p *JSONParser) Get(key string) (any, bool) {
	if v, ok := p.Data[key]; ok {
		return v, true
	} else {
		return nil, false
	}
}
