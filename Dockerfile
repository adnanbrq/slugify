FROM golang:alpine

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o slugify ./cmd/main.go

CMD [ "./slugify" ]