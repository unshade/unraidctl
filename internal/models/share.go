package models

import "fmt"

type Share struct {
	Name string `json:"name"`
	Free int    `json:"free"`
	Used int    `json:"used"`
}

func (s *Share) String() string {
	return fmt.Sprintf("Name: %s | Free: %d | Used: %d", s.Name, s.Free, s.Used)
}

type ListSharesModel struct {
	Shares []Share `json:"shares"`
}

func (m *ListSharesModel) String() string {
	result := "Shares:\n"
	for _, share := range m.Shares {
		result += share.String() + "\n"
	}
	return result
}
