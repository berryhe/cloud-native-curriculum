FROM golang:1.17.2 as builder
WORKDIR /go/src/github.com/berryhe/cloud-native-curriculum/week_01
COPY . /go/src/github.com/berryhe/cloud-native-curriculum/week_01
RUN go mod init github.com/duty-os/sync-table-validator && \
    GOPROXY="https://goproxy.io,direct" go mod tidy && \
    GOOS=linux CGO_ENABLED="0" go build -ldflags="-s -w" -o week1


FROM alpine:latest as prod
WORKDIR /app/
COPY --from=builder /go/src/github.com/berryhe/cloud-native-curriculum/week_01/week1 /app/week1
CMD ["/app/week1"]