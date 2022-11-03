package ctl

const COMPOSE_VERSION = "3.8"
const REDIS_IMAGE = "redis"
const REDIS_IMAGE_TAG = "6.2.3"

func NewDockerComposeConfig(stack Stack) ([]byte, error) {

	return []byte{}, nil
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
