#!/bin/bash -eu
set -o pipefail

# Install mcpcurl
mkdir -p ~/.local/bin
pushd /tmp
    git clone --depth 1 https://github.com/github/github-mcp-server.git
    cd github-mcp-server

    go build -o ~/.local/bin/mcpcurl cmd/mcpcurl/main.go
popd

# Install yq
mkdir -p "${HOME}/.local/bin"
curl -sSfL -o - https://github.com/mikefarah/yq/releases/download/v4.45.1/yq_linux_amd64.tar.gz | \
    tar -zxf - -O "./yq_linux_amd64" > "${HOME}/.local/bin/yq"
chmod +x "${HOME}/.local/bin/yq"

# Install SonarScanner CLI
mkdir -p "${HOME}/.local/"
SONAR_VERSION='7.1.0.4889'
curl -sSfL \
    -o 'sonar-scanner-cli-linux-x64.zip' \
    --output-dir '/tmp' \
    "https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-${SONAR_VERSION}-linux-x64.zip"
unzip '/tmp/sonar-scanner-cli-linux-x64.zip' -d "${HOME}/.local/"
mv "${HOME}/.local/sonar-scanner-${SONAR_VERSION}-linux-x64" "${HOME}/.local/sonar-scanner"
echo 'export PATH=$PATH:~/.local/sonar-scanner/bin' >> ~/.bashrc

#export SONAR_TOKEN=
#sonar-scanner \
#    -Dsonar.host.url=http://127.0.0.1:9000 \
#    -Dsonar.projectKey=sonarqube-mcp-server \
#    -Dsonar.sources=.
