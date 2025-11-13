package models

type ListContainersModel struct {
	Docker struct {
		Containers []struct {
			ID    string `json:"id"`
			Image string `json:"image"`
			State string `json:"state"`
		} `json:"containers"`
	} `json:"docker"`
}

type StartContainerModel struct {
	Docker struct {
		Start struct {
			ID string `json:"id"`
		} `json:"start"`
	} `json:"docker"`
}


type StopContainerModel struct {
	Docker struct {
		Stop struct {
			ID string `json:"id"`
		} `json:"stop"`
	} `json:"docker"`
}
