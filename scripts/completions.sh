#!/bin/sh
set -e
mkdir completions
for sh in bash zsh fish; do
	go run main.go completion "$sh" >"completions/voiceflow.$sh"
done