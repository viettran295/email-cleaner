FROM golang:1.21-alpine3.19 AS build 
WORKDIR /app
COPY . /app
RUN go mod download && go mod verify
RUN go build -o /app/email-cleaner

FROM alpine:latest 
WORKDIR /app
COPY credentials.json /app
COPY token.json /app
COPY --from=build /app/email-cleaner /app
ENTRYPOINT [ "/app/email-cleaner" ]