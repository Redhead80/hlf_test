package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ApplicationConfiguration struct {
	Port int `yaml:"port"`
}

type ClientConfiguration struct {
	ConnectionProfile string `yaml:"connection_profile"`
	UserName          string `yaml:"user_name"`
	Organization      string `yaml:"organization"`
	Channel           string `yaml:"channel"`
}

type Configuration struct {
	Application ApplicationConfiguration `yaml:"application"`
	Client      ClientConfiguration      `yaml:"client"`
}

func LoadConfigurationFromFile(configFile string) (Configuration, error) {

	configFileContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		return Configuration{}, err
	}

	var configuration Configuration
	err = yaml.Unmarshal(configFileContent, &configuration)
	return configuration, err
}
