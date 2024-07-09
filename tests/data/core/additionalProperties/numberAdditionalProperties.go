// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"

type NumberAdditionalProperties struct {
	// Name corresponds to the JSON schema field "name".
	Name *string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	AdditionalProperties map[string]float64
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NumberAdditionalProperties) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain NumberAdditionalProperties
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw[""]; !ok || v == nil {
		plain.AdditionalProperties = map[string]float64{}
	}
	*j = NumberAdditionalProperties(plain)
	return nil
}