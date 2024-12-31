#!/bin/bash

set -e
set -x
set -v
rm -rf completions

mkdir completions
for sh in bash zsh fish; do
	go run main.go completion "$sh" >"completions/voiceflow.$sh"
done