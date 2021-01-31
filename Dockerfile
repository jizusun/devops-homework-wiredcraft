# https://docs.docker.com/develop/develop-images/multistage-build/
FROM golang:alpine AS builder
WORKDIR /app
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux 
# https://dev.to/ivan/go-build-a-minimal-docker-image-in-just-three-steps-514i
COPY src/go.mod .
COPY src/go.sum .
RUN go mod download
COPY src/ .
RUN ls -la .
RUN go build -o hugoctl .

FROM klakegg/hugo:alpine AS hugo
RUN apk --no-cache add fortune git
WORKDIR /website
COPY . .
COPY --from=builder /app/hugoctl ./hugoctl
RUN ["./hugoctl", "dev"]  

# https://duske.me/modern-multiarch-builds-with-docker/

# FROM golang:alpine AS builder
# RUN mkdir /app
# ADD . /app/
# WORKDIR /app
# RUN go build -o report .

# FROM busybox
# RUN mkdir /app
# WORKDIR /app
# COPY --from=builder /app/report .
# CMD ["./report"]