package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"jsonparser/internal/cli"
	"jsonparser/internal/constants"
	"jsonparser/internal/entity"
)

func main() {
	if len(os.Args) > 2 {
		panic(constants.ErrorTooManyArgs)
	}

	p, err := cli.GetFilePath()
	if err != nil {
		panic(err)
	}

	parser := entity.JSONParser{}

	if _, err = parser.DeserializeJSON(p); err != nil {
		panic(err)
	}

	var done bool

	for !done {
		var key string
		fmt.Printf(constants.MessageEnterAKey)
		fmt.Scanln(&key)

		if v, ok := parser.Get(key); !ok {
			slog.Error(constants.ErrorKeyNotFound, "key", key)
		} else {
			slog.Info("value found", "key", v)
		}

		fmt.Printf(constants.MessageTryAgain)

		var answer string
		fmt.Scanln(&answer)

		if strings.ToLower(answer) != "y" {
			done = true
		}

	}
}
