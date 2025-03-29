package utils

import (
	"io"

	"github.com/xavidop/voiceflow-cli/internal/global"
)

// SafeClose closes an io.Closer and logs any error
func SafeClose(c io.Closer) {
	if err := c.Close(); err != nil {
		global.Log.Errorf("Error closing resource: %v", err)
	}
}
