FROM golang:1.20 as builder

WORKDIR /app

COPY . .
RUN go get

WORKDIR ./main
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

WORKDIR /app
RUN go build -o name ./main/.

FROM golang
COPY --from=builder /app/name /app/
COPY --from=builder /app/docs /app/

EXPOSE 3000

ENTRYPOINT [ "/app/name" ]



#docker build . -t Container-name:latest
#docker run -p 3000:3000 Container-name:latest

