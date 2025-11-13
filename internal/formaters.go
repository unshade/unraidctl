package internal

import (
	"encoding/json"

	"go.yaml.in/yaml/v2"
)

type OutputFormater[T any] interface {
	Format(input T) (string, error)
}

type YamlOutputFormater struct{}

var _ OutputFormater[any] = (*YamlOutputFormater)(nil)

func (f *YamlOutputFormater) Format(input any) (string, error) {
	yamled, err := yaml.Marshal(input)
	if err != nil {
		return "", err
	}
	return string(yamled), nil
}

type JsonOutputFormater struct{}

var _ OutputFormater[any] = (*JsonOutputFormater)(nil)

func (f *JsonOutputFormater) Format(input any) (string, error) {
	jsoned, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		return "", err
	}
	return string(jsoned), nil
}

type ArrayOutputFormater struct{}

var _ OutputFormater[ArrayInput] = (*ArrayOutputFormater)(nil)

type ArrayInput struct {
	Header []string
	Content [][]string
}
func (f *ArrayOutputFormater) Format(input ArrayInput) (string, error) {
	jsoned, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		return "", err
	}
	return string(jsoned), nil
}
