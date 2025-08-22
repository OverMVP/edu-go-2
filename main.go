package main

import (
	"fmt"
	"os"
	"strings"

	"jsonparser/internal/cli"
	"jsonparser/internal/constants"
	"jsonparser/internal/entity"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println(constants.TOO_MANY_ARGS)
		os.Exit(1)
	}

	p, err := cli.GetFilePath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	parser := entity.JSONParser{}

	if _, err = parser.DeserializeJSON(p); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var done bool

	for !done {
		var key string
		fmt.Printf(constants.TITLE_ENTER_A_KEY)
		fmt.Scanln(&key)

		if v, ok := parser.Get(key); !ok {
			fmt.Println(constants.KEY_NOT_FOUND)
		} else {
			fmt.Printf("%v\n", v)
		}

		fmt.Printf(constants.TITLE_TRY_AGAIN)

		var answer string
		fmt.Scanln(&answer)

		if strings.ToLower(answer) != "y" {
			done = true
		}

	}
}
