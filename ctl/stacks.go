package ctl

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

var StackTypes = []string{"docker-compose"}

type Stack struct {
	Name        string      `yaml:"name"`
	RedisAddr   string      `yaml:"redisAddr"`
	ManageRedis bool        `yaml:"manageRedis"`
	Kind        string      `yaml:"kind"`
	Apps        []AppConfig `yaml:"apps"`
}

type AppConfig struct {
	ID             string   `yaml:"id"`
	InboundTopics  []string `yaml:"inboundTopics,omitempty"`
	OutboundTopics []string `yaml:"outboundTopics,omitempty"`
	Ports          []string `yaml:"ports,omitempty"`
	Image          Image    `yaml:"image"`
	ExtraArgs      []string `yaml:"extra_args"`
}

type Image struct {
	Name string `yaml:"name"`
	Tag  string `yaml:"tag"`
}

// Generate a new, empty configuration
func NewStack(kind, name, redisAddr string) (*Stack, error) {
	s := new(Stack)
	s.Name = name
	s.Kind = kind
	s.RedisAddr = redisAddr

	return s, nil
}

func NewStackWithPrompt() (*Stack, error) {
	var (
		stackType string
		stackName string
		redisAddr string
	)
	s, err := NewStack(stackType, stackName, redisAddr)
	if err != nil {
		return new(Stack), err
	}

	// Set name
	stackNamePrompt := promptui.Prompt{
		Label: "Stack name",
	}
	s.Name, err = stackNamePrompt.Run()
	if err != nil {
		return s, err
	}

	// Set kind
	stackTypePrompt := promptui.Select{
		Label: "Stack type",
		Items: StackTypes,
	}
	_, s.Kind, err = stackTypePrompt.Run()

	if err != nil {
		return s, err
	}

	switch s.Kind {
	case "docker-compose":
		fmt.Println("docker-compose")
	default:
		fmt.Println("default")
	}

	// Set manageRedis
	manageRedisSelection := promptui.Select{
		Label: "Manage redis for you?",
		Items: []bool{true, false},
	}
	_, res, err := manageRedisSelection.Run()
	if err != nil {
		return s, err
	}

	manageRedis, err := strconv.ParseBool(res)
	if err != nil {
		return s, err
	}

	s.ManageRedis = manageRedis

	// Set redis
	redisAddrPrompt := promptui.Prompt{
		Label: "Redis address",
	}
	s.RedisAddr, err = redisAddrPrompt.Run()

	if err != nil {
		return s, err
	}

	return s, nil
}
