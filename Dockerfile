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


FROM scratch
#FROM alpine:latest AS run

WORKDIR /src

COPY --from=build /src/main .
COPY --from=build /src/config config
COPY --from=build /src/database database

#RUN apk add ffmpeg
#RUN apk add postgresql18 postgresql18-contrib

EXPOSE 8081

CMD [ "./main" ]