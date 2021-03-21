FROM golang:alpine
WORKDIR /
ADD . /
RUN go build -o app
ENTRYPOINT ./app
EXPOSE 8000
