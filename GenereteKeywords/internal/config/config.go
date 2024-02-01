package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	CountKeywords     int    `json:"count_keywords"`
	NameOutFileResult string `json:"name_out_file_result"`
	NameInputDataFile string `json:"name_input_data_file"`
}

func GetConfig() *Config {
	// Read the JSON file
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		//fmt.Println("Error reading JSON file:", err)
		return nil
	}

	// Parse JSON data into a Configuration struct
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		//fmt.Println("Error parsing JSON:", err)
		return nil
	}
	return &config
}
