package ctl

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

const COMPOSE_VERSION = "3.8"
const REDIS_IMAGE = "redis"
const REDIS_IMAGE_TAG = "6.2.3"

type ComposeFile struct {
	Version  string                    `yaml:"version"`
	Services map[string]ComposeService `yaml:"services"`
}

func (c *ComposeFile) ToYaml() ([]byte, error) {
	return yaml.Marshal(c)
}

type ComposeService struct {
	ContainerName string   `yaml:"container_name"`
	Image         string   `yaml:"image"`
	Build         string   `yaml:"build,omitempty"`
	Command       []string `yaml:"command"`
	DependsOn     []string `yaml:"depends_on,omitempty"`
	EnvFile       []string `yaml:"env_file,omitempty"`
}

func NewDockerComposeConfig(stack Stack) (*ComposeFile, error) {
	config := new(ComposeFile)
	config.Services = make(map[string]ComposeService)
	prefix := stack.Name
	config.Version = COMPOSE_VERSION

	if stack.ManageRedis {
		redisSvc := ComposeService{
			ContainerName: fmt.Sprintf("%s-%s", prefix, "redis"),
			Image:         fmt.Sprintf("%s:%s", REDIS_IMAGE, REDIS_IMAGE_TAG),
		}
		config.Services[fmt.Sprintf("%s-redis", prefix)] = redisSvc
	}
	for _, svc := range stack.Apps {
		var command []string
		for _, topic := range svc.InboundTopics {
			command = append(command, "-inbound")
			command = append(command, topic)
		}

		for _, topic := range svc.OutboundTopics {
			command = append(command, "-outbound")
			command = append(command, topic)
		}

		command = append(command, "-redis")
		command = append(command, fmt.Sprintf(stack.RedisAddr))
		command = append(command, svc.ExtraArgs...)

		composeSvc := ComposeService{
			ContainerName: fmt.Sprintf("%s-%s", prefix, svc.ID),
			Image:         fmt.Sprintf("%s:%s", svc.Image.Name, svc.Image.Tag),
			Command:       command,
		}
		config.Services[svc.ID] = composeSvc
	}
	return config, nil
}

/*
version: '3.8'
services:
  bytebot:
    container_name: fraqbot-gateway
    image: ghcr.io/bytebot-chat/gateway-discord:latest
    command:
      - '-id'
      - 'discord'
      - '-inbound'
      - 'discord-inbound'
      - '-outbound'
      - 'discord-outbound'
      - '-t'
      - '$TOKEN'
      - '-redis'
      - 'redis:6379'
    depends_on:
      - redis
    env_file:
      - .env
  party-pack:
    container_name: fraqbot-party-pack
    build: ghcr.io/bytebot-chat/bytbot-party-pack:latest
    command:
      - '-discord-inbound'
      - 'discord-inbound'
      - '-discord-outbound'
      - 'discord-outbound'
    depends_on:
      - bytebot
  fraqbot-py:
    container_name: fraqbot-py
    build: app/
    depends_on:
      - bytebot
    env_file:
      - .env
  redis:
    container_name: fraqbot-redis
    image: redis:6.2.3
    ports:
      - '127.0.0.1:6379:6379'
*/
