FROM golang:1.15 as builder

ENV GO111MODULE on
ENV SRC_DIR $GOPATH/src/git.iwanhae.kr/wan/jupynetes

COPY . ${SRC_DIR}
WORKDIR ${SRC_DIR}
RUN go install -mod=vendor

FROM golang:1.15
WORKDIR /bin
COPY --from=builder /go/bin/jupynetes /bin

CMD ["jupynetes"]