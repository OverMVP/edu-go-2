package cli

import (
	"errors"
	"flag"
	"fmt"
	"path/filepath"

	"jsonparser/internal/constants"
)

func GetFilePath() (p string, e error) {
	flag.StringVar(&p, "path", "", "path to a file")
	flag.Parse()

	if p == "" {
		return p, errors.New(constants.NO_PATH)
	}

	extension := filepath.Ext(p)

	if extension != constants.EXT_JSON {
		fmt.Println()
		return p, fmt.Errorf(constants.EXPECTED_JSON, extension)
	}

	return p, nil
}
