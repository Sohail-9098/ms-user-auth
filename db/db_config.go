package db

import (
	"os"
	"path/filepath"

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

var (
	configFileName string = filepath.Join(getWorkDir(), "config/config.yaml")
)

func loadConfig() (DbConfig, error) {
	var config DbConfig
	file := util.OpenFile(configFileName)
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err := decoder.Decode(&config)
	util.HandleError("error umarshal file: ", err)
	return config, nil
}

func getWorkDir() string {
	cwd, err := os.Getwd()
	util.HandleError("error get working dir: ", err)
	return cwd
}
