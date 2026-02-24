FROM golang:alpine AS build

WORKDIR /src

COPY go.mod .
COPY go.sum .

COPY . .

ENV GOCACHE=/root/.cache/go-build

RUN --mount=type=cache,target="/root/.cache/go-build" \
    GOOS=linux go build -o main cmd/*.go

FROM alpine:latest AS run

WORKDIR /src

COPY --from=build /src/main .
COPY --from=build /src/config/. .

EXPOSE 8081


CMD [ "./main" ]