package Trial

import (
	"encoding/json"
	"os"
	"text/template"
)

// {
// "input": "welcome to the world of <<params.provider>>. You will be working on <<params.project>>.",
// "params": [
// {
// "name": "provider",
// "value": "kubernetes"
// },
// {
// "name": "project",
// "value": "project-1"
// }
// ]
// }
type JsonInput struct {
	Input  string `json:"input"`
	Params []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"params"`
}

func GetJsonInput(input string) {
	// Define the JSON data
	jsonData := `{
		"input": "welcome to the world of <<params.provider>>. You will be working on <<params.project>>.",
		"params": [
			{
				"name": "provider",
				"value": "kubernetes"
			},
			{
				"name": "project",
				"value": "project-1"
			}
		]
	}`

	// Unmarshal the JSON data into a struct
	var data JsonInput
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		panic(err)
	}

	// Define the Go template
	tmpl := template.Must(template.New("template").Parse(data.Input))

	// Execute the template with the JSON data
	err = tmpl.Execute(os.Stdout, data.Params)
	if err != nil {
		panic(err)
	}
}
