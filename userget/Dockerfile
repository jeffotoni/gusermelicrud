FROM golang:1.17 as builder
WORKDIR /go/src/userget
COPY . . 
RUN cp userget /go/bin/userget

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/userget /
CMD ["/userget"]
