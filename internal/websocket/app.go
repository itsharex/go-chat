package main

import (
	"go-chat/config"
	"go-chat/internal/provider"
	"go-chat/internal/websocket/internal/process"
)

type Provider struct {
	Config    *config.Config
	Server    provider.WebsocketServer
	Coroutine *process.Server
}
