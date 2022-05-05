FROM golang:1.18.1-alpine
RUN mkdir /app
RUN apk add git
COPY . /app
WORKDIR /app
RUN go build -o server .
CMD [ "/app/server" ]