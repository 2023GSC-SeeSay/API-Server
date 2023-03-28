# Use an official Golang runtime as a parent image
FROM golang:1.20.0

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

RUN go mod download
# install python and required packages for g2pk
RUN apt-get update && apt-get install -y python3 python3-pip && pip3 install --upgrade pip && pip3 install g2pk

# Build the Go app
RUN go build -o api-server .
# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./api-server"]
