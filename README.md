[![CircleCI](https://circleci.com/gh/jmatsu/artifact-transfer/tree/master.svg?style=svg)](https://circleci.com/gh/jmatsu/artifact-transfer/tree/master)

# transart

transart - Transfer artifacts from services to the single service

```bash
transart [-f <configuration file>] <command> [command options...]
```

## Installation

Download the binary directly

```bash
curl -sL "https://raw.githubusercontent.com/jmatsu/transart/master/install.bash" | bash
curl -sL "https://raw.githubusercontent.com/jmatsu/transart/master/install.bash" | VERSION=<...> bash
```

Or build on your local

```bash
dep ensure
go build .
```

## Getting started

```bash
# Create the configuration file
transart init --save-dir ".transart"

# Configure GitHub Release as a *source* service
transart add github-release \
    --source \
    --username jmatsu \
    --reponame transart \
    --api-token-name GITHUB_TOKEN \
    --file-name-pattern ".*\.tar$" # tar files only

# Configure Local file system as a *destination* service
transart add local \
    --destination \
    --path tmp # use ./tmp directory

export GITHUB_TOKEN=...

# Download artifacts from GitHub Release and copy only tar files of them to 'tmp' directory
transart transfer

# Download artifacts only
transart download

# Upload artifacts only
transart upload
```

### Supported services

- CircleCI (source only)
- GitHub Release (destination only)
- Local file system

### Configurations

#### Basic structure

```
version: 1
save_dir: <path>
source:
  locations:
  - <location config>
destination:
  location: <location config>
```

#### Location config

ref. [](../config/option_key.go)

*CircleCI*

Only source is supported

```yaml
type: circleci
vcs-type: <github|bitbucket>
username: <the username of the project> # case sensitive
reponame: <the name of the project> # case sensitive

# optional
branch: <the name of the branch>
api-token-name: <environment name> #CIRCLECI_TOKEN is used by default
```

*GitHub Release*

Only distination is supported

```yaml
type: github-release
strategy: <draft|create|draft_or_create>
username: <the username of the repository>
reponame: <the name of the respoitory>

# optional
api-token-name: <environment name> #GITHUB_TOKEN is used by default
```

*Local file system*

```yaml
type: local
path: <path>

# optional
file-name-pattern: <regexp pattern>
```

### Bash/Zsh completion

```bash
// For Bash
eval $(transart --init-completion bash)
transart --init-completion bash >> ~/.bashrc


// For Zsh
eval $(transart --init-completion zsh)
transart --init-completion zsh >> ~/.zshrc
```

## LICENSE

Under MIT License. See [LICENSE](./LICENSE)
