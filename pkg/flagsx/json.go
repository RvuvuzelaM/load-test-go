package flagsx

import "encoding/json"

type JSONFlag struct {
	Target interface{}
}

func (f *JSONFlag) String() string {
	b, err := json.Marshal(f.Target)
	if err != nil {
		return "failed to marshal object"
	}
	return string(b)
}

func (f *JSONFlag) Set(v string) error {
	return json.Unmarshal([]byte(v), f.Target)
}

// Type is only used in help text
func (f *JSONFlag) Type() string {
	return "json"
}
