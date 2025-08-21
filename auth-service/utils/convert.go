package utils

import "fmt"

func ToStringSlice(raw interface{}) ([]string, error) {
	rawSlice, ok := raw.([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid type: expected []interface{}, got %T", raw)
	}

	result := make([]string, 0, len(rawSlice))
	for _, item := range rawSlice {
		str, ok := item.(string)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected string, got %T", item)
		}
		result = append(result, str)
	}

	return result, nil
}