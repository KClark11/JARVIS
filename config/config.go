package config

import (
	"encoding/json"
	"fmt"       // used to print errors
	"io/ioutil" //used to help read config.json file.
)

var (
	Token     string        //strores value of token from config.json file.
	BotPrefix string        //stores value of BotPrefix from confi.json.
	config    *configStruct //stores value of extracted from config.json.
)

type configStruct struct {
	Token     string `json : "Token"`
	BotPrefix string `json : "BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(string(file))
	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil

}
