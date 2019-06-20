FROM golang:alpine as go
WORKDIR /go/src/app
COPY ./hello.go .
RUN go build -v


FROM scratch
COPY --from=go /go/src/app/app /
CMD [ "/app" ]