# build stage
FROM golang:1.14-alpine AS build

# env
RUN mkdir -p /app
WORKDIR /app

# tools
RUN apk add --no-cache git
# build src
COPY go.mod .
RUN go mod download

# app src
COPY . .
RUN go build -o /bin/app

# result stage
FROM alpine:latest
COPY --from=build /bin/app /bin/app
RUN mkdir -p /views
COPY ./views ./views
ENTRYPOINT ["/bin/app", "--port", "8080"]