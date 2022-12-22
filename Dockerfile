FROM golang:1.18

## node & generator
RUN apt update \
 && apt install -y nodejs npm \
 && npm install -g openapi-generator

## jq
RUN apt install -y jq

## for lambda development
ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /usr/bin/aws-lambda-rie
RUN chmod +x /usr/bin/aws-lambda-rie

## hot reload
RUN go install github.com/cosmtrek/air@latest

## linter
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1

WORKDIR /workspace

CMD ["air", "-c", ".air.toml"]
