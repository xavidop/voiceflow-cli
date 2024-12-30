# Install

You can use `voiceflow-cli` by installing a pre-compiled binary (in several ways), using Docker, or compiling it from source. In the below sections, you can find the steps for each approach.

## Install a pre-compiled binary

### homebrew tap
Install the Voiceflow CLI:
```sh
brew install xavidop/tap/voiceflow
```

### scoop

```powershell
scoop bucket add voiceflow https://github.com/xavidop/scoop-bucket.git
scoop install voiceflow
```

### chocolatey

```powershell
choco install voiceflow
```

### nix

#### nixpkgs

```bash
nix-env -iA voiceflow
```

!!! info
    The [package in nixpkgs](https://github.com/NixOS/nixpkgs/blob/main/pkgs/tools/misc/voiceflow/default.nix)
    might be slightly outdated, as it is not updated automatically.
    Use our NUR to always get the latest updates.

#### nur

First, you'll need to add our [NUR](https://github.com/xavidop/nur) to your nix configuration.
You can follow the guides
[here](https://github.com/nix-community/NUR#installation).

Once you do that, you can install the packages.

```nix
{ pkgs, lib, ... }: {
    home.packages = with pkgs; [
    nur.repos.xavidop.voiceflow
    ];
}
```

### deb, rpm and apk packages

Download the `.deb`, `.rpm` or `.apk` packages from the [OSS releases page][releases] and install them with the appropriate tools.

### go install

```sh
go install github.com/xavidop/voiceflow-cli@latest
```

### bash script

```sh
curl -sfL https://voiceflow.xavidop.me/static/run | bash
```

#### Additional Options
You can also set the `VERSION` variable to specify
a version instead of using latest.

You can also pass flags and args to voiceflow-cli:

```bash
curl -sfL https://voiceflow.xavidop.me/static/run |
    VERSION=__VERSION__ bash -s -- version
```

!!! tip
    This script does not install anything, it just downloads, verifies and
    runs voiceflow-cli.
    Its purpose is to be used within scripts and CIs.

### manually

Download the pre-compiled binaries from the [releases page][releases] and copy them to the desired location.


## Verifying the artifacts

### binaries

All artifacts are checksummed, and the checksum file is signed with [cosign][].

1. Download the files you want along with the `checksums.txt`, `checksum.txt.pem`, and `checksums.txt.sig` files from the [releases][releases] page:
    ```sh
    wget https://github.com/xavidop/voiceflow-cli/releases/download/__VERSION__/checksums.txt
    wget https://github.com/xavidop/voiceflow-cli/releases/download/__VERSION__/checksums.txt.sig
    wget https://github.com/xavidop/voiceflow-cli/releases/download/__VERSION__/checksums.txt.pem
    ```
1. Verify the signature:
    ```sh
    COSIGN_EXPERIMENTAL=1 cosign verify-blob \
    --cert checksums.txt.pem \
    --signature checksums.txt.sig \
    checksums.txt
    ```
1. If the signature is valid, you can then verify the SHA256 sums match with the downloaded binary:
    ```sh
    sha256sum --ignore-missing -c checksums.txt
    ```

### docker images

Our Docker images are signed with [cosign][].

Verify the signatures:

```sh
COSIGN_EXPERIMENTAL=1 cosign verify xavidop/voiceflow
```

!!! info
    The `.pem` and `.sig` files are the image `name:tag`, replacing `/` and `:` with `-`.

## Running with Docker

You can also use `voiceflow-cli` within a Docker container.
To do that, you'll need to execute something more-or-less like the examples below.

Registries:

- [`xavidop/voiceflow`](https://hub.docker.com/r/xavidop/voiceflow)
- [`ghcr.io/xavidop/voiceflow`](https://github.com/xavidop/voiceflow-cli/pkgs/container/voiceflow)

Example usage:

```sh
docker run --rm \
    xavidop/voiceflow voiceflow version
```

Note that the image will almost always have the last stable Go version.

If you need other packages and dependencies, you are encouraged to keep your own image. You can
always use voiceflow-cli's [own Dockerfile][dockerfile] as a starting point and iterate on that.

[dockerfile]: https://github.com/xavidop/voiceflow-cli/blob/main/Dockerfile
[releases]: https://github.com/xavidop/voiceflow-cli/releases
[cosign]: https://github.com/sigstore/cosign

## Compiling from source

Here you have two options:

If you want to contribute to the project, please follow the
steps on our [contributing guide](/community/contributing/).

If you just want to build from source for whatever reason, follow these steps:

**clone:**

```sh
git clone https://github.com/xavidop/voiceflow-cli
cd voiceflow-cli
```

**get the dependencies:**

```sh
go mod tidy
```

**build:**

```sh
go build -o voiceflow .
```

**verify that it works:**

```sh
./voiceflow version
```
