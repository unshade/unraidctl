package internal

type OutputFormat string

const (
	JSON OutputFormat = "json"
	YAML OutputFormat = "yaml"
	TEXT OutputFormat = "text"
)

func OutputFormaterSwitcher(output OutputFormat) OutputFormater {
	var formater OutputFormater
	switch output {
	case "json":
		formater = &JsonOutputFormater{}
	case "yaml":
		formater = &YamlOutputFormater{}
	case "text":
		formater = &TextOutputFormater{}
	default:
		formater = &TextOutputFormater{}
	}
	return formater
}
