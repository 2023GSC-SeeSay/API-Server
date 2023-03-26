# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

RUN go mod download

# Build the Go app
RUN go build -o api-server .
RUN gsutil cp gs://firebase-cred/seesay-firebase-adminsdk-clpnw-faf918ab9f.json /app/credentials.json
# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./api-server", "-creds", "credentials.json", "-port", "8080:8080"]
