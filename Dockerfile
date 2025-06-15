FROM golang:1.24.3-bookworm AS build

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,target=.,Z \
    CGO_ENABLED=0 go build -o /bin/sonarqube-mcp-server ./cmd/sonarqube-mcp-server/main.go

FROM gcr.io/distroless/static-debian12

COPY --from=build /bin/sonarqube-mcp-server /

ENTRYPOINT ["/sonarqube-mcp-server"]
