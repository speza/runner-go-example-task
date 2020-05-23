FROM golang:alpine AS build_base

RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

FROM build_base AS build_go

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' .

CMD ["./runner-go-example-task"]
EXPOSE 5300
