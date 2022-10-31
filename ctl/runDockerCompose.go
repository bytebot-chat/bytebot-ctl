package ctl

import "gopkg.in/yaml.v2"

const COMPOSE_VERSION = "3.8"
const REDIS_IMAGE = "redis"
const REDIS_IMAGE_TAG = "6.2.3"

// DockerComposeConfig represents the unique options required to generate a full
// docker-compose stack
type DockerComposeConfig struct {
	CreateRedis bool             `yaml:"-"`
	NamePrefix  string           `yaml:"-"`
	Services    []ComposeService `yaml:"services"`
	Version     string           `yaml:"version"`
}

// ComposeService represents a service in Docker Compose
type ComposeService struct {
	ContainerName string   `yaml:"container_name"`
	Image         string   `yaml:"image"`
	Command       []string `yaml:"command"`
	DependsOn     string   `yaml:"depends_on"`
	EnvFile       string   `yaml:"env_file"`
	Build         string   `yaml:"build"`
	Ports         []string `yaml:"ports"`
}

// NewDockerComposeConfig generates a DockerComposeConfig struct suitable for use
// In generating a complete docker-compose.yaml file using NewDockerComposeFile()
func NewDockerComposeConfig(namePrefix string, createRedis bool) DockerComposeConfig {
	return DockerComposeConfig{
		CreateRedis: true,
		NamePrefix:  namePrefix,
	}
}

// NewDockerComposeFile takes a more limited DockerComposeConfig and expands it into
// A full, working, opinionated docker-compose file represeted as a byte slice ([]byte)
func NewDockerComposeFile(config DockerComposeConfig) ([]byte, error) {
	config.Version = COMPOSE_VERSION
	if config.CreateRedis {
		redisConfig := ComposeService{
			ContainerName: "redis-" + config.NamePrefix,
			Image:         REDIS_IMAGE + ":" + REDIS_IMAGE_TAG,
			Ports:         []string{"127.0.0.1:6379:6379"},
		}

		config.Services = append(config.Services, redisConfig)
	}
	return yaml.Marshal(config)
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
