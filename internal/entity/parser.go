package entity

import (
	"encoding/json"
	"os"
)

type (
	ParsedData map[string]any

	JSONParser struct {
		Data ParsedData
	}
)

func (p *JSONParser) DeserializeJSON(path string) (ParsedData, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(raw, &p.Data); err != nil {
		return nil, err
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
