FROM golang:1.21-alpine

WORKDIR /app

COPY main.go .

ENV PORT=8080

EXPOSE ${PORT}

CMD [ "go", "run", "main.go" ]