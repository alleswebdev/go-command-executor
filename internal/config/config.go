package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port     int
	Name     string
	Commands map[string]Command
}

type Command struct {
	Name    string
	Command string
	Start   string
	Stop    string
	Restart string
}

func GetAppConfig() Config {
	cfg := Config{}

	viper.SetConfigName("values")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Errorf("fatal unmarshal config file: %w", err))
	}

	return cfg
}

//func all() {
//	for _, group := range cfg.CommandGroups {
//		cmd := exec.Command("/bin/bash", "-c", group.Command)
//
//		var stdoutBuf, stderrBuf bytes.Buffer
//		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
//		cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
//
//		err := cmd.Start()
//
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//
//		cmd.Wait()
//		fmt.Println(stdoutBuf.String())
//	}
//}
