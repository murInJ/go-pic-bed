FROM golang:1.22-alpine AS build

COPY handler/ /go/src/handler/
COPY middleware/ /go/src/middleware/
COPY router/ /go/src/router/
COPY service/ /go/src/service/
COPY utils/ /go/src/utils/
COPY /docs /go/src/docs
COPY go.mod go.sum *.go /go/src/

WORKDIR "/go/src/"
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
ENV GOOS=linux 
ENV GOARCH=amd64
RUN go mod tidy
RUN go build -o pic-bed


FROM alpine:latest

RUN mkdir "/app"
COPY --from=build /go/src/pic-bed /app/pic-bed
COPY config.json /app/config.json
COPY data/ /app/data/


RUN chmod +x /app/pic-bed

EXPOSE 18001
WORKDIR "/app"
ENTRYPOINT ["/app/pic-bed"]