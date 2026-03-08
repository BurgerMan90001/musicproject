FROM golang:alpine AS build

WORKDIR /src

COPY . .

ENV GOCACHE=/root/.cache/go-build

RUN --mount=type=cache,target="/root/.cache/go-build" \
    GOOS=linux go build -o main cmd/*.go

FROM alpine:latest AS run

WORKDIR /src

COPY --from=build /src/main .
COPY --from=build /src/config config
COPY --from=build /src/schema schema


RUN apk add postgresql18 postgresql18-contrib

CMD [ "./main" ]