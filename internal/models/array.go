package models

type ShowArrayModel struct {
	Array struct {
		State string `json:"state"`
	} `json:"array"`
}

type MutateArrayModel struct {
	Id string `json:"id"`
}
