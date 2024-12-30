# Frequently Asked Questions (FAQ)

## How does it work?

`voiceflow-cli` has three main purposes:

1. Make the interaction with your Voiceflow agents from your laptop or your continuous integration pipelines easier than ever
2. Create testing tools that will help users build their Vocieflow agent

## Who is `voiceflow-cli` for?

`voiceflow-cli` is primarily for software engineering teams who are currently using Voiceflow. It is recommended for machine learning engineers that usually work with LLMs, STT, TTS, NLU and NLP technologies.

## What kind of machines/containers do I need for the `voiceflow-cli`?

You'll need either: a bare-metal host (your own, AWS i3.metal or Equinix Metal) or a VM that supports nested virtualisation such as those provided by Google Cloud, Azure, AWS, DigitalOcean, etc. or a Linux or Windows container.

## When will Jenkins, GitLab CI, BitBucket Pipeline Runners, Drone or Azure DevOps be supported?

For the current phase, we're targeting GitHub Actions because it has fine-grained access controls and the ability to schedule exactly one build to a runner. The other CI systems will be available soon.

That said, if you're using these tools within your organisation, we'd like to hear from you.
So feel free to reach out to us if you feel `voiceflow-cli` would be a good fit for your team.

Feel free to contact us at: [xavierportillaedo@gmail.com](mailto:xavierportillaedo@gmail.com)

## What kind of access is required in my Voiceflow project?

Refer to the Authentication page [here](/overview/authentication)

## Can voiceflow-cli be used on public repos?

Yes, `voiceflow-cli` can be used on public and private repos.

## What's in the Container image and how is it built?

The Container image contains uses `alpine:latest` and the `voiceflow-cli` installed on it.

The image is built automatically using GitHub Actions and is available on a container registry.

## Is ARM64 supported?

Yes, `voiceflow-cli` is built to run on both Intel/AMD and ARM64 hosts. This includes a Raspberry Pi 4B, AWS Graviton, Oracle Cloud ARM instances and potentially any other ARM64 instances that support virtualisation.

## Are Windows or macOS supported?

Yes, in addition to Linux, Windows and macOS are also supported platforms for `voiceflow-cli` at this time on a AMD64 or ARM64 architecture.

## Is `voiceflow-cli` free and open-source?

`voiceflow-cli` is an open source tool, however, it interacts with Voiceflow, so a Voiceflow account is required.

The website and documentation are available on GitHub and we plan to release some open source tools in the future for voiceflow customers.