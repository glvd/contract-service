package contract

import "encoding/json"

// VideoMessage ...
type VideoMessage struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Version string `json:"version"`
}

// DecodeMessage ...
func DecodeMessage(s string) (m VideoMessage, e error) {
	e = json.Unmarshal([]byte(s), &m)
	return m, e
}

// DecodeMessages ...
func DecodeMessages(s []string) (m []VideoMessage, e error) {
	for _, ss := range s {
		message, e := DecodeMessage(ss)
		if e != nil {
			return nil, e
		}
		m = append(m, message)
	}
	return m, nil
}

// Encode ...
func (m VideoMessage) Encode() (s string, e error) {
	bys, e := json.Marshal(m)
	if e != nil {
		return "", e
	}
	return string(bys), e
}

// MustJSON ...
func MustJSON(s string, e error) string {
	if e == nil {
		return s
	}
	panic(e)
}
