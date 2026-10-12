FROM ghcr.io/ez4bk/devops/go-compiler AS builder

WORKDIR /workspace
COPY go.mod /workspace/go.mod
COPY go.sum /workspace/go.sum
RUN go mod download

COPY . /workspace
RUN CGO_ENABLED=0 go build -ldflags='-s -w' -o /workspace/app ./cmd/...


FROM alpine

WORKDIR /workspace
COPY config /workspace/config
COPY schemas /workspace/schemas
COPY --from=builder /workspace/app /workspace/app

EXPOSE 8899/tcp
EXPOSE 11211/tcp
ENTRYPOINT ["/workspace/app"]
