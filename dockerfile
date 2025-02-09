FROM golang:1.22.5
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY app ./
COPY proto ./
COPY idl ./
COPY main.go ./
COPY Makefile ./
RUN go install github.com/moulid15/BetOnBet@latest
RUN go get github.com/moulid15/BetOnBet
RUN CGO_ENABLED=0 go build -o main
CMD ["./main"]