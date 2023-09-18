FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY main.go .

RUN go build main.go

FROM scratch AS runner

WORKDIR /app

COPY --from=builder /app/main /app

EXPOSE 8080

CMD [ "./main" ]
