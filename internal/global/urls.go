package global

import "fmt"

// resolveSubdomain returns the formatted subdomain string.
// If an override is provided it takes precedence, otherwise the global is used.
// A non-empty value is returned with a leading dot (e.g. ".staging").
func resolveSubdomain(override string) string {
	subdomain := VoiceflowSubdomain
	if override != "" {
		subdomain = override
	}
	if subdomain != "" {
		subdomain = "." + subdomain
	}
	return subdomain
}

// GetAPIBaseURL returns the base URL for the Voiceflow creator-api service.
// Priority: VoiceflowAPIURL > VoiceflowSubdomain > default.
// Example defaults:
//
//	https://api.voiceflow.com
//	https://api.staging.voiceflow.com  (subdomain = "staging")
func GetAPIBaseURL(subdomainOverride string) string {
	if VoiceflowAPIURL != "" {
		return VoiceflowAPIURL
	}
	subdomain := resolveSubdomain(subdomainOverride)
	return fmt.Sprintf("https://api%s.voiceflow.com", subdomain)
}

// GetRuntimeBaseURL returns the base URL for the Voiceflow general-runtime service.
// Priority: VoiceflowRuntimeURL > VoiceflowSubdomain > default.
func GetRuntimeBaseURL(subdomainOverride string) string {
	if VoiceflowRuntimeURL != "" {
		return VoiceflowRuntimeURL
	}
	subdomain := resolveSubdomain(subdomainOverride)
	return fmt.Sprintf("https://general-runtime%s.voiceflow.com", subdomain)
}

// GetAnalyticsBaseURL returns the base URL for the Voiceflow analytics-api service.
// Priority: VoiceflowAnalyticsURL > VoiceflowSubdomain > default.
func GetAnalyticsBaseURL() string {
	if VoiceflowAnalyticsURL != "" {
		return VoiceflowAnalyticsURL
	}
	subdomain := resolveSubdomain("")
	return fmt.Sprintf("https://analytics-api%s.voiceflow.com", subdomain)
}
