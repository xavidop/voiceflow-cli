#!/bin/sh
set -e
mkdir manpages
go run . man | gzip -c -9 >manpages/voiceflow.1.gz