package models

type ShowArrayModel struct {
	Array struct {
		State string `json:"state"`
	} `json:"array"`
}

func (m *ShowArrayModel) String() string {
	return "Array State: " + m.Array.State
}

type MutateArrayModel struct {
	Id string `json:"id"`
}

func (m *MutateArrayModel) String() string {
	return "Array ID: " + m.Id
}
