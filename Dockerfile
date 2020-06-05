FROM golang:latest
WORKDIR /app
ADD . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app


# Stage Runtime Applications
FROM alpine:latest

# Setting timezone
ENV TZ=Asia/Jakarta
RUN apk add -U tzdata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /home/app

COPY --from=0 /app/app .

EXPOSE 8080 8081

ENTRYPOINT [ "./app" ]
