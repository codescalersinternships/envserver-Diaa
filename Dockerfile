FROM golang:1.20 AS build_stage

WORKDIR /app


COPY . ./


RUN CGO_ENABLED=0 go build -o bin/app cmd/main.go


EXPOSE 8080

FROM alpine:latest 

WORKDIR /app/

COPY --from=build_stage ./app/bin/app ./

CMD [ "./app","-p","8080" ]
