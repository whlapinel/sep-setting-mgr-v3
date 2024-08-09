
FROM golang:latest

WORKDIR /app

COPY main .

RUN chmod +x ./main

EXPOSE 1323

CMD ["./main"]
