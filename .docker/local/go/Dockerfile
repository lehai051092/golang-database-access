
FROM golang:1.19.4-alpine as base

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache gcc musl-dev make git bash openssh

WORKDIR /var/www/app
