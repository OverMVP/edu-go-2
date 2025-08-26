package cli

import (
	"flag"
	"fmt"
	"path/filepath"

	"jsonparser/internal/constants"
)

func GetFilePath() (p string, err error) {
	flag.StringVar(&p, "path", "", "path to a file")
	flag.Parse()

	if p == "" {
		return p, fmt.Errorf(constants.ErrorNoPathGiven)
	}

	extension := filepath.Ext(p)

	if extension != constants.ExtensionJSON {
		return p, fmt.Errorf(constants.ErrorExpectedJSON, extension)
	}

	return p, nil
}
