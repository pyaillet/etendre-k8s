FROM golang:1.11 as builder

WORKDIR /
COPY main.go /

RUN go mod init server
RUN CGO_ENABLED=0 go build .

FROM scratch

COPY --from=builder server /server
COPY template.html cat.gif /

ENTRYPOINT [ "/server" ]
EXPOSE 8080

