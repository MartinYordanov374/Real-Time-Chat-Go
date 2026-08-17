FROM golang:1.25.0

WORKDIR /backend

COPY ../go.mod ../go.sum ./

RUN go mod download

COPY . .

CMD ["go", "run", "./Backend/main.go"]
