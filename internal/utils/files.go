package utils

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/xavidop/voiceflow-cli/internal/types/tests"
	"gopkg.in/yaml.v3"
)

func CheckIfFileExists(file string) error {

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return err
	}

	return nil
}

func WriteFile(b []byte, file string) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	err = os.WriteFile(file, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetRelativeFilePathFromParentFile(parentFile string, file string) string {
	base := filepath.Dir(parentFile)

	if !filepath.IsAbs(file) {
		return path.Join(base, file)
	} else {
		return file
	}
}

func LoadSuitesFromPath(path string) ([]tests.Suite, error) {
	var suites []tests.Suite

	// Convert the provided path to an absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve path %s: %w", path, err)
	}

	files, err := os.ReadDir(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", absPath, err)
	}

	for _, file := range files {

		filePath := filepath.Join(absPath, file.Name())
		// Check if the file is a YAML file
		if !file.IsDir() && filepath.Ext(filePath) == ".yaml" {
			data, err := os.ReadFile(filePath)
			if err != nil {
				return []tests.Suite{}, fmt.Errorf("failed to read file %s: %v", filePath, err)
			}

			var suite tests.Suite
			if err := yaml.Unmarshal(data, &suite); err != nil {
				return []tests.Suite{}, fmt.Errorf("failed to parse YAML file %s: %v", filePath, err)
			}

			// Resolve relative paths for test files
			for i, test := range suite.Tests {
				testAbsPath, err := filepath.Abs(filepath.Join(filepath.Dir(filePath), test.File))
				if err != nil {
					return []tests.Suite{}, fmt.Errorf("failed to resolve test file path %s: %v", test.File, err)
				}
				suite.Tests[i].File = testAbsPath
			}
			suites = append(suites, suite)
		}
	}
	return suites, err
}

func LoadTestFromPath(path string) (tests.Test, error) {
	// Check if the file is a YAML file
	if filepath.Ext(path) == ".yaml" {
		data, err := os.ReadFile(path)
		if err != nil {
			return tests.Test{}, fmt.Errorf("failed to read file %s: %v", path, err)
		}

		var test tests.Test
		if err := yaml.Unmarshal(data, &test); err != nil {
			return tests.Test{}, fmt.Errorf("failed to parse YAML file %s: %v", path, err)
		}

		return test, nil
	} else {
		return tests.Test{}, fmt.Errorf("file is not a YAML file %s", path)
	}
}
