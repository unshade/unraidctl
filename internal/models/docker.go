package models

import (
	"fmt"
	"strings"
)

type ListContainersModel struct {
	Docker struct {
		Containers []struct {
			ID    string `json:"id"`
			Image string `json:"image"`
			State string `json:"state"`
		} `json:"containers"`
	} `json:"docker"`
}

func (m *ListContainersModel) String() string {
	builder := strings.Builder{}
	for _, container := range m.Docker.Containers {
		idParts := strings.Split(container.ID, ":")
		compactID := idParts[len(idParts)-1]

		if len(compactID) > 12 {
			compactID = compactID[:12]
		}

		builder.WriteString(fmt.Sprintf("ID: %s | Image: %s | State: %s\n", compactID, container.Image, container.State))
	}
	return builder.String()
}

type StartContainerModel struct {
	Docker struct {
		Start struct {
			ID string `json:"id"`
		} `json:"start"`
	} `json:"docker"`
}

func (m *StartContainerModel) String() string {
	return "Started Container ID: " + m.Docker.Start.ID
}

type StopContainerModel struct {
	Docker struct {
		Stop struct {
			ID string `json:"id"`
		} `json:"stop"`
	} `json:"docker"`
}

func (m *StopContainerModel) String() string {
	return "Stopped Container ID: " + m.Docker.Stop.ID
}
