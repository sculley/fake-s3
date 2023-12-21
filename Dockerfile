# Compile stage
FROM golang:1 AS build-env

RUN mkdir /app
COPY . /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

# Run build process
RUN CGO_ENABLED=0 go build -ldflags "-X main.build=${VCS_REF}" -o /app/cmd/fake-s3 /app/cmd/fake-s3

FROM alpine:latest

RUN mkdir /etc/fake-s3
RUN mkdir /etc/fake-s3/conf

COPY ./conf /etc/fake-s3/conf

WORKDIR /etc/fake-s3
COPY --from=build-env /app/cmd/fake-s3/fake-s3 /bin/fake-s3

VOLUME [ "/etc/fake-s3/conf" ]

EXPOSE 8080

CMD ["/bin/fake-s3"]