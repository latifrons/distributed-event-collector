FROM golang:1.22-bullseye AS builder

RUN useradd -m builder
WORKDIR /src
RUN go env -w GOPRIVATE=github.com/atomeight/*

ARG ACCESS_TOKEN_USR="nothing"
ARG ACCESS_TOKEN_PWD="nothing"

RUN printf "machine github.com\n\
    login ${ACCESS_TOKEN_USR}\n\
    password ${ACCESS_TOKEN_PWD}\n\
    \n\
    machine api.github.com\n\
    login ${ACCESS_TOKEN_USR}\n\
    password ${ACCESS_TOKEN_PWD}\n"\
    >> /root/.netrc
RUN chmod 600 /root/.netrc

ADD ./go.mod ./go.sum ./
RUN go mod download

COPY . .

RUN chown -R builder:builder /src
USER builder

RUN go build -a -o main .

# Copy OG into basic alpine image
FROM ubuntu:24.04

RUN apt-get update && apt-get install -y curl=8.5.0-2ubuntu10.3 tzdata=2024a-2ubuntu1 && rm -rf /var/lib/apt/lists/*
RUN useradd -m exchange

WORKDIR /www
COPY --from=builder --chown=exchange:exchange /src/data/config ./data/config/
COPY --from=builder --chown=exchange:exchange /src/main .

RUN chown -R exchange:exchange .

USER exchange

ENTRYPOINT ["./main"]