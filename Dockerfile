FROM golang:1.18

## for lambda development
ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /usr/bin/aws-lambda-rie
RUN chmod +x /usr/bin/aws-lambda-rie

## hot reload
RUN go install github.com/cosmtrek/air@latest

WORKDIR /workspace

CMD ["air", "-c", ".air.toml"]
