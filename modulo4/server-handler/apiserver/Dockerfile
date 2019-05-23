#################################################
# Dockerfile distroless
#################################################
FROM golang:1.12.4 as builder
WORKDIR /go/src/apiserver
ENV GO111MODULE=on
COPY . .
#RUN GOOS=linux go build -ldflags="-s -w" -o apiserver main.go
RUN cp apiserver /go/bin/apiserver
RUN ls -lh
#RUN go install -v ./...
# path/go/bin

############################
# STEP 2 build a small image
############################
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/apiserver /
CMD ["/apiserver"]