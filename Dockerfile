FROM alpine:latest

COPY voiceflow_*.apk /tmp/
RUN apk add --no-cache --allow-untrusted /tmp/voiceflow_*.apk