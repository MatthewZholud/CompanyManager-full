FROM golang:alpine

RUN mkdir /app 
ADD . /app/
WORKDIR /app/cmd

RUN go build -o main .
RUN adduser -S -D -H -h /app appuser

USER appuser
CMD ["./main"]
