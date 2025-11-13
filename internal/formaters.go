package internal

import (
	"encoding/json"
	"fmt"

	"go.yaml.in/yaml/v2"
)

type OutputFormater interface {
	Format(input any) (string, error)
}

type YamlOutputFormater struct{}

var _ OutputFormater = (*YamlOutputFormater)(nil)

func (f *YamlOutputFormater) Format(input any) (string, error) {
	yamled, err := yaml.Marshal(input)
	if err != nil {
		return "", err
	}
	return string(yamled), nil
}

type JsonOutputFormater struct{}

var _ OutputFormater = (*JsonOutputFormater)(nil)

func (f *JsonOutputFormater) Format(input any) (string, error) {
	jsoned, err := json.MarshalIndent(input, "", " ")
	if err != nil {
		return "", err
	}
	return string(jsoned), nil
}

type TextOutputFormater struct{}

var _ OutputFormater = (*TextOutputFormater)(nil)

func (f *TextOutputFormater) Format(input any) (string, error) {
	if stringer, ok := input.(fmt.Stringer); ok {
		return stringer.String(), nil
	}
	return "", fmt.Errorf("input does not implement fmt.Stringer")
}

func PrintFormat(stringed string, err error) {
	if err != nil {
		fmt.Printf("Error formatting output: %v\n", err)
		return
	}
	fmt.Println(stringed)
}
