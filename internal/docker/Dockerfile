# internal/docker/Dockerfile
FROM gcc:latest

# Install coreutils to get the `timeout` command (it's part of coreutils)
RUN apt-get update && apt-get install -y coreutils

# Set the working directory inside the container
WORKDIR /app
