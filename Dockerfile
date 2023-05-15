# Start from golang base image
FROM golang:alpine as builder

RUN mkdir /app

##set direktori utama
WORKDIR /app

##copy seluruh file ke app
ADD . /app

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["/app/main"]