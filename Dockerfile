FROM golang:latest
# Set the Current Working Directory inside the container
WORKDIR /app/
ENV PORT 8080
EXPOSE 8080
COPY . . 
RUN go mod download
RUN go build -o omdb .
CMD ["./omdb"]
