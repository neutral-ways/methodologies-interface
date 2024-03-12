# Build stage
FROM golang:1.18-alpine3.15 AS builder
WORKDIR /app
COPY . .

RUN apk update
# RUN apk add --virtual build-dependencies
RUN apk add libc-dev
RUN apk add make
RUN apk add gcc
RUN apk add bash
RUN apk add git

#RUN go get -u github.com/swaggo/swag/cmd/swag
#RUN make doc

RUN make build

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/corezero-api .
COPY app.env .
COPY db ./db
COPY domain/workers/mailer/templates ./domain/workers/mailer/templates

EXPOSE 8080
CMD [ "/app/corezero-api" ]
ENTRYPOINT [ "/app/corezero-api" ]