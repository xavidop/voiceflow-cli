package global

import "github.com/sirupsen/logrus"

var Log logrus.Logger

var VoiceflowAPIKey string
var OpenAIAPIKey string
var OpenAIBaseURL string
var VoiceflowSubdomain string

// Custom base URL overrides for each Voiceflow service.
// When set, these take priority over VoiceflowSubdomain.
var VoiceflowAPIURL string       // Overrides the Voiceflow API (creator-api) base URL
var VoiceflowRuntimeURL string   // Overrides the general-runtime base URL
var VoiceflowAnalyticsURL string // Overrides the analytics-api base URL

var VersionString string
var Output string
var Verbose bool
var SkipUpdate bool
var ShowTesterMessages bool
