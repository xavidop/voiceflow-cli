/*
Copyright © 2023 Xavier Portilla Edo

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"os"

	"github.com/xavidop/voiceflow-cli/cmd"
	"github.com/xavidop/voiceflow-cli/internal/server"
)

func main() {
	// Register server flags with cobra root command
	cmd.RegisterServerFlags()

	// Check if it's a version or help command first
	if len(os.Args) > 1 && (os.Args[1] == "version" || os.Args[1] == "help" || os.Args[1] == "--help" || os.Args[1] == "-h") {
		cmd.Execute()
		return
	}

	// Check if server mode is enabled via flag or env var
	serverEnabled := os.Getenv("SERVER") == "true"
	if !serverEnabled && len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			if arg == "--server" || arg == "-s" {
				serverEnabled = true
				break
			}
		}
	}

	if serverEnabled {
		// Get subdomain from flag or env var
			subdomain := os.Getenv("VOICEFLOW_SUBDOMAIN")
			if subdomain == "" {
				for i, arg := range os.Args {
					if (arg == "--voiceflow-subdomain" || arg == "-b") && i+1 < len(os.Args) {
						subdomain = os.Args[i+1]
						break
					}
				}
			}
			if subdomain != "" {
				os.Setenv("VOICEFLOW_SUBDOMAIN", subdomain)
			}

			srv := server.NewServer()
			// Get port from flag or env var
			port := os.Getenv("PORT")
			if port == "" {
				for i, arg := range os.Args {
					if (arg == "--port" || arg == "-p") && i+1 < len(os.Args) {
						port = os.Args[i+1]
						break
					}
				}
			}
			if err := srv.Run(port); err != nil {
				panic(err)
			}
			return
	}

	// Default CLI mode
	cmd.Execute()
}
