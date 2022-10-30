package ctl

import (
	"github.com/manifoldco/promptui"
)

var StackTypes = []string{"local", "docker", "docker-compose"}

type Stack struct {
	Name      string      `yaml:"name"`
	RedisAddr string      `yaml:"redisAddr"`
	Kind      string      `yaml:"kind"`
	Apps      []AppConfig `yaml:"apps"`
}

type AppConfig struct {
	ID             string   `yaml:"id"`
	InboundTopics  []string `yaml:"inboundTopics"`
	OutboundTopics []string `yaml:"outboundTopics"`
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
