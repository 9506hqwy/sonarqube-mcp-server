#!/bin/bash
set -euo pipefail

VERSION=v9.9

BASE_DIR=$(dirname "$(dirname "$0")")
TMP_DIR=/tmp
BASE_DIR="${BASE_DIR}/pkg/sonarqube"

curl -fsSL \
    -o "${TMP_DIR}/openapi.yml" \
     "https://github.com/9506hqwy/openapi-spec-sonarqube-v1/raw/refs/heads/${VERSION}/openapi.yml"

function capitalize() {
    VALUE="$1"
    # shellcheck disable=SC2206
    IFS=_ ARR=(${VALUE})

    CAP_ARR=()
    for V in "${ARR[@]}"
    do
        CAP_ARR+=("${V^}")
    done

    IFS='' CAP="${CAP_ARR[*]}"

    echo "${CAP}"
}

function groupname() {
    OP_PATH="$1"
    echo "${OP_PATH}" | cut -d '/' -f 3
}

function apiname() {
    OP_PATH="$1"
    echo "${OP_PATH}" | cut -d '/' -f 4
}

function filename() {
    OP_PATH="$1"
    OP_GROUP=$(groupname "${OP_PATH}")
    echo "${BASE_DIR}/${OP_GROUP}.go"
}

function toolname() {
    OP_PATH="$1"
    OP_GROUP=$(groupname "${OP_PATH}")
    OP_API=$(apiname "${OP_PATH}")
    echo "${OP_GROUP,,}_${OP_API,,}"
}

function write-preceding() {
    OP_PATH="$1"

    FILE_PATH=$(filename "${OP_PATH}")
    cat > "${FILE_PATH}" <<EOF
package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)
EOF
}

function write-register-func() {
    OP_PATH="$1"
    DEFINITION="$2"

    DESCRIPTION=$(yq -r '.description' <<<"${DEFINITION}")
    TOOL_NAME=$(toolname "${OP_PATH}")
    API_NAME=$(capitalize "${TOOL_NAME}")
    NUM_PARAMETER=$(yq -r '.parameters | length' <<<"${DEFINITION}")

    echo "| ${TOOL_NAME} | ${DESCRIPTION} |"

    FILE_PATH=$(filename "${OP_PATH}")
    # signature
    cat >> "${FILE_PATH}" <<EOF

func register${API_NAME}(s *server.MCPServer) {
	tool := mcp.NewTool("${TOOL_NAME}",
		mcp.WithDescription("${DESCRIPTION//\"/\\\"}"),
EOF

    # parameters
    HANDER_NAME="${API_NAME,}Handler"
    if [[ ${NUM_PARAMETER} -ne 0 ]]; then
        HANDER_NAME="mcp.NewTypedToolHandler(${HANDER_NAME})"
        cat >> "${FILE_PATH}" <<EOF
		mcp.WithInputSchema[client.Api${API_NAME}Params](),
EOF
    fi

    # complete
    cat >> "${FILE_PATH}" <<EOF
	)

	s.AddTool(tool, ${HANDER_NAME})
}
EOF
}

function write-handler-func() {
    OP_PATH="$1"
    DEFINITION="$2"

    TOOL_NAME=$(toolname "${OP_PATH}")
    API_NAME=$(capitalize "${TOOL_NAME}")
    NUM_PARAMETER=$(yq -r '.parameters | length' <<<"${DEFINITION}")

    PARSE_STMT=", params client.Api${API_NAME}Params"
    PARAM_ARG="&params,"
    if [[ ${NUM_PARAMETER} -eq 0 ]]; then
        PARSE_STMT=""
        PARAM_ARG=""
    fi

    FILE_PATH=$(filename "${OP_PATH}")
    cat >> "${FILE_PATH}" <<EOF

func ${API_NAME,}Handler(ctx context.Context, request mcp.CallToolRequest ${PARSE_STMT}) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.Api${API_NAME}(ctx, ${PARAM_ARG} authorizationHeader))
}
EOF
}

yq -r '.paths | keys | .[]' "${TMP_DIR}/openapi.yml" | while read -r OP_PATH
do
    FILE_PATH=$(filename "${OP_PATH}")
    rm -rf "${FILE_PATH}"
done

yq -r '.paths | keys | .[]' "${TMP_DIR}/openapi.yml" | while read -r OP_PATH
do
    write-preceding "${OP_PATH}"
done

yq -r '.paths | keys | .[]' "${TMP_DIR}/openapi.yml" | while read -r OP_PATH
do
    OP=$(yq ".paths.\"${OP_PATH}\"" "${TMP_DIR}/openapi.yml")

    OP_POST=$(yq '.post' <<<"${OP}")
    if [[ "$OP_POST" != 'null' ]]; then
        write-register-func "${OP_PATH}" "${OP_POST}"
        write-handler-func "${OP_PATH}" "${OP_POST}"
    fi

    OP_GET=$(yq '.get' <<<"${OP}")
    if [[ "$OP_GET" != 'null' ]]; then
        write-register-func "${OP_PATH}" "${OP_GET}"
        write-handler-func "${OP_PATH}" "${OP_GET}"
    fi
done

# write tools.go
TOOLS_PATH="${BASE_DIR}/tools.go"
cat > "${TOOLS_PATH}" <<EOF
package sonarqube

import (
	"github.com/mark3labs/mcp-go/server"
)

func RegisterTools(s *server.MCPServer, readonly bool) {
EOF

yq -r '.paths | keys | .[]' "${TMP_DIR}/openapi.yml" | while read -r OP_PATH
do
    TOOL_NAME=$(toolname "${OP_PATH}")
    API_NAME=$(capitalize "${TOOL_NAME}")

    OP=$(yq ".paths.\"${OP_PATH}\"" "${TMP_DIR}/openapi.yml")

    OP_POST=$(yq '.post' <<<"${OP}")
    if [[ "$OP_POST" != 'null' ]]; then
        echo "if !readonly { register${API_NAME}(s) }" >> "${TOOLS_PATH}"
    fi

    OP_GET=$(yq '.get' <<<"${OP}")
    if [[ "$OP_GET" != 'null' ]]; then
        echo "register${API_NAME}(s)" >> "${TOOLS_PATH}"
    fi
done

cat >> "${TOOLS_PATH}" <<EOF
}
EOF

# format
go fmt ./...
