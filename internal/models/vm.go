package models

type ListVMsModel struct {
	VMs struct {
		ID      string `json:"id"`
		Domains []struct {
			Name  string `json:"name"`
			State string `json:"state"`
			Id    string `json:"id"`
		} `json:"domains"`
	} `json:"vms"`
}

type StartVMModel struct {
	Id string `json:"id"`
}

type StopVMModel struct {
	Id string `json:"id"`
}
