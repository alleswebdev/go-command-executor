package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

type config struct {
	Port   int
	Name   string
	Groups map[string]Group
}

type Group struct {
	Name    string
	Command string
	Start   string
	Stop    string
	Restart string
}

func main() {
	cfg := config{}

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
		fmt.Println(err)
	}

	fmt.Println(cfg)

	for _, group := range cfg.Groups {
		cmd := exec.Command("/bin/bash", "-c", group.Command)

		var stdoutBuf, stderrBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
		cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

		err = cmd.Start()

		if err != nil {
			fmt.Println(err.Error())
		}

		cmd.Wait()
		fmt.Println(stdoutBuf.String())
	}
}
