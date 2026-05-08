FROM golang:alpine AS build

WORKDIR /src

ENV GOCACHE=/go-cache
ENV GOMODCACHE=/gomod-cache

COPY ./go.* ./

RUN --mount=type=cache,target=/gomod-cache \
      go mod download

COPY ./ ./

RUN --mount=type=cache,target=/gomod-cache \
   --mount=type=cache,target=/go-cache \
   GOOS=linux go build -o main cmd/musicproject/main.go


FROM alpine:latest

WORKDIR /src

COPY --from=build /src/main .
COPY --from=build /src/.env.dev .env.dev
COPY --from=build /src/config.dev.yml config.dev.yml


# RUN apk add --no-cache \
#     ca-certificates
#RUN apk add ffmpeg


EXPOSE 8081

CMD [ "./main" ]