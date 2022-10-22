FROM golang:alpine 

# RUN apk update && apk add --no-cache git
WORKDIR /app
COPY /Praktikum .
RUN go mod tidy
RUN go build -o bin 
ENTRYPOINT [ "/app/bin" ]
EXPOSE 8080