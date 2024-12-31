#!/bin/bash

set -e

rm -rf completions
mkdir completions
go version
for sh in bash zsh fish; do
	go run main.go completion "$sh" >"completions/voiceflow.$sh"
done