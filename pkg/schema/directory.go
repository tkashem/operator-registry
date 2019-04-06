package schema

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ValidateManifestInDirectory(manifestDir string) error {
	return validate(manifestDir)
}

func validate(directory string) error {
	err := filepath.Walk(directory, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".yaml") {
			return nil
		}

		fmt.Printf("validate %s\n", path)
		if validateResourceInFile(path, f, err) != nil {
			return err
		}

		return nil
	})
	return err
}

func validateResourceInFile(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	exampleFileReader, err := os.Open(path)
	if err != nil {
		return err
	}
	defer exampleFileReader.Close()

	// return ValidateResource( exampleFileReader )
	return nil
}
