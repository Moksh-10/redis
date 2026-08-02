FROM golang:1.25

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

EXPOSE 7379

CMD ["go", "run", "main.go"]