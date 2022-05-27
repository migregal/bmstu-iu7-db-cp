FROM golang:alpine as builder
LABEL maintainer="Gregory @migregal Mironov"

RUN mkdir -p /src/backend
WORKDIR /src/backend

ADD . .

RUN cd ./cmd/cube && go build -o /bin/app


FROM alpine
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

WORKDIR /bin/

COPY --from=builder /bin/app .

ENV GIN_MODE=release

ENV ADDR=0.0.0.0

EXPOSE 3000

CMD exec /bin/app
