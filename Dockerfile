FROM golang:1.12 AS build
COPY . /app
WORKDIR /app
RUN go build -o main .

FROM golang:alpine AS runtime
WORKDIR /app
COPY --from=build /app/main ./
CMD ["./main"]
EXPOSE 5300
