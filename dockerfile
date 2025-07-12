# Specifies a parent image
FROM golang:1.24.4

# Creates an app directory to hold your app’s source code
WORKDIR /app

# Copies everything from your root directory into /app
COPY . .

# Installs Go dependencies
RUN go mod download

# Builds your app with optional configuration
RUN go build -o examen-tecnico-stori 

# Tells Docker which network port your container listens on
EXPOSE 8080

# Specifies the executable command that runs when the container starts
CMD [ "./examen-tecnico-stori"]

# Set environment variable with api key
ENV SENDGRID_API_KEY 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
