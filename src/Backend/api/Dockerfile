FROM golang:1.24.4

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

ENV PORT=8080
EXPOSE 8080

CMD ["go", "run", "."]