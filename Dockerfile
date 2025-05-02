# Build Stage
FROM golang:1.21 AS build

WORKDIR /SPADE
COPY . ./
RUN go build -o /bin/server ./usecases/hypnogram/cmd/server/
RUN go build -o /bin/analyst ./usecases/hypnogram/cmd/analyst/
RUN go build -o /bin/user ./usecases/hypnogram/cmd/user/
