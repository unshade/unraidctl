package models

import (
	"fmt"
	"strings"
)

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

func (m *ListVMsModel) String() string {
	builder := strings.Builder{}
	for _, vm := range m.VMs.Domains {
		idParts := strings.Split(vm.Id, ":")
		compactID := idParts[len(idParts)-1]

		builder.WriteString(fmt.Sprintf("ID: %s | Name: %s | State: %s\n", compactID, vm.Name, vm.State))
	}
	return builder.String()
}

type StartVMModel struct {
	Id string `json:"id"`
}

func (m *StartVMModel) String() string {
	return "Started VM ID: " + m.Id
}

type StopVMModel struct {
	Id string `json:"id"`
}

func (m *StopVMModel) String() string {
	return "Stopped VM ID: " + m.Id
}
