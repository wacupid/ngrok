package auth

import (
	"ngrok/log"
	"strings"
)

type Auth interface {
	Auth(token string) bool
}

func New() Auth {
	return &MyAuth{}
}

type MyAuth struct{}

var secretPath = "/etc/ngrok-secrets"

func SetSecretPath(path string) {
	secretPath = path
}

func (this *MyAuth) Auth(token string) bool {
	fields := strings.Split(token, ":")
	if len(fields) != 2 {
		log.Info("wrong format of token")
		return false
	}
	username := fields[0]
	password := fields[1]

	parser := NewParser(secretPath)
	if err := parser.Parse(); err != nil {
		log.Warn("Parser:", err)
		return false
	}
	if val, ok := parser.Tokens[username]; ok {
		if val == password {
			return true
		}
	}
	return false
}