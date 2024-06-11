package db

import (
	"github.com/go-yaml/yaml"
	"github.com/sohail-9098/ms-user-auth/util"
)

type DbConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	SslMode  string `yaml:"sslmode"`
	Db       string `yaml:"db"`
}

const configFilePath string = "/usr/local/bin/config.yaml"

func LoadConfig() (DbConfig, error) {
	var config DbConfig
	file, err := util.OpenFile(configFilePath)
	if err != nil{
		return DbConfig{}, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err!=nil{
		return DbConfig{}, err
	}

	return config, nil
}
