FROM alpine:3.5
MAINTAINER Gaurav CHoudhary

COPY ./Golang-Microservice /app/Golang-Microservice
RUN chmod +x /app/Golang-Microservice

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT /app/Golang-Microservice
