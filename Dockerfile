FROM golang:latest
# Set the Current Working Directory inside the container
WORKDIR /app/
EXPOSE 5000
COPY . . 

RUN go mod download
RUN go build -o omdb .
CMD ["./omdb"]
