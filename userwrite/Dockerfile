FROM golang:1.17 as builder
WORKDIR /go/src/userwrite
COPY . . 
RUN cp userwrite /go/bin/userwrite

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/userwrite /
CMD ["/userwrite"]
