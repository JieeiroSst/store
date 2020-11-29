FROM golang:1.13-alpine
ENV LANG='ja_JP.utf8'
ENV GO111MODULE='on'

ARG POSTGRES_DB
ARG POSTGRES_USER
ARG POSTGRES_PASSWORD
ARG POSTGRES_PORT
ENV POSTGRES_DB $POSTGRES_DB
ENV POSTGRES_USER $POSTGRES_USER
ENV POSTGRES_PASSWORD $POSTGRES_PASSWORD
ENV POSTGRES_PORT $POSTGRES_PORT

WORKDIR /go/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./api /go/app

CMD realize start