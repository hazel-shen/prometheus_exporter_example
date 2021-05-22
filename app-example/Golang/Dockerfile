FROM golang:alpine
WORKDIR /
COPY . /
RUN go build -o app
ENTRYPOINT ./app
EXPOSE 8000
