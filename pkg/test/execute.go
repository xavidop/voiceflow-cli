package test

import (
	"log"

	"github.com/google/uuid"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/utils"
)

func ExecuteSuite(suitesPath string) error {

	// Define the user ID
	userID := uuid.New().String()

	// Load all suites from the path
	suites, err := utils.LoadSuitesFromPath(suitesPath)
	if err != nil {
		log.Fatalf("Error loading suites: %v", err)
	}

	// Iterate over each suite and its tests
	for _, suite := range suites {
		global.Log.Infof("Suite: %s\nDescription: %s\nEnvironment: %s", suite.Name, suite.Description, suite.EnvironmentName)
		global.Log.Infof("Running Tests:")

		for _, testFile := range suite.Tests {
			test, err := utils.LoadTestFromPath(testFile.File)
			if err != nil {
				log.Fatalf("Error loading test: %v", err)
			}
			err = runTest(suite.EnvironmentName, userID, test)
			if err != nil {
				log.Fatalf("Error running test: %v", err)
			}
		}
	}
	return nil
}
