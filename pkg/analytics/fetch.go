package analytics

import (
	"fmt"

	"github.com/xavidop/voiceflow-cli/pkg/voiceflow"
)

func Fetch(agentID, outputFile, startTime, endTime string, limit int, analyticsToFetch []string) error {

	startTimeDate, endTimeDate, limitInt, err := voiceflow.ParseFilters(startTime, endTime, limit)
	if err != nil {
		return fmt.Errorf("failed to parse filters: %w", err)
	}

	analytics, err := voiceflow.FetchAnalytics(agentID, startTimeDate, endTimeDate, limitInt, analyticsToFetch)
	if err != nil {
		return fmt.Errorf("failed to fetch analytics %s: %w", agentID, err)
	}

	if outputFile != "" {
		err = voiceflow.SaveAnalytics(analytics, outputFile)
		if err != nil {
			return fmt.Errorf("failed to save analytics with agentID %s: %w", agentID, err)
		}
	}

	fmt.Printf("%s\n", analytics)

	return nil
}
