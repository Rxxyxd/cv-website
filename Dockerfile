FROM golang:1.18.1-alpine
EXPOSE 8080/tcp
RUN mkdir /app
RUN apk add git
COPY . /app
WORKDIR /app
RUN go build -o server .
CMD [ "/app/server" ]