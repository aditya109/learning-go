package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/**
1.	Open `Questions` directory
2. 	Iterate through all yaml files starting with `Question#..`
3.	Retrieve Issue YAML
4.	Customize Create Issue GitHub API with YAML
*/

type JSONConfig struct {
	Issue Issue `json:"issue"`
}

type Issue struct {
	Question Question `json:"question"`
	Metadata Metadata `json:"metadata"`
}

type Question struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Metadata struct {
	Labels    []string `json:"labels"`
	Milestone string   `json:"milestone"`
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func getJSON() (jsonConfig JSONConfig) {
	jsonFile, err := os.Open("D:\\Projects\\go-workspace\\src\\github.com\\aditya109\\json-reader\\Question#1.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &jsonConfig)

	return
}
func main() {
	getJSON()
}
