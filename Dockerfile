FROM golang:1.15 as builder

ENV GO111MODULE on

COPY . /app
WORKDIR /app
RUN go build -o jupynetes main.go
RUN chmod +x jupynetes

FROM ubuntu
COPY --from=builder /app/jupynetes /usr/bin
CMD ["/usr/bin/jupynetes"]