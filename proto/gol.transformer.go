package gol

import (
	"github.com/cjp2600/gol/core"
)

func (m *Job) GetUnmarshalBody() (map[string]interface{}, error) {
	var resp = make(map[string]interface{})
	for k, m := range m.Body {
		value, err := core.GetAnyType(m)
		if err != nil {
			return nil, err
		}
		resp[k] = value
	}
	return resp, nil
}

func (m *Job) GetUnmarshalHeader() (map[string]string, error) {
	var resp = make(map[string]string)
	for k, m := range m.Header {
		value, err := core.GetAnyType(m)
		if err != nil {
			return nil, err
		}
		resp[k] = value.(string)
	}
	return resp, nil
}
