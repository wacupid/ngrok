package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"ngrok/log"
)

type Parser struct {
	Path    string
	Content string
	Tokens  map[string]string
}

func NewParser(path string) *Parser {
	return &Parser{
		Path:   path,
		Tokens: make(map[string]string),
	}
}

func (this *Parser) Parse() error {
	file, err := os.Open(this.Path)
	if err != nil {
		err = fmt.Errorf("Failed to read configuration file %s: %v", this.Path, err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		toParse := strings.Split(text, "#")[1]
		log.Info("wacupid toParse: ========> " + toParse)
		fields := strings.Fields(toParse)
		if len(fields) != 2 {
			continue
		}
		log.Info("wacupid fields0: ========> " + fields[0])
		log.Info("wacupid fields1: ========> " + fields[1])
		this.Tokens[fields[0]] = fields[1]
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}