FROM golang:1.20 AS build_stage

WORKDIR /app


COPY . ./


RUN CGO_ENABLED=0 go build -o bin/app cmd/main.go


EXPOSE 8080

FROM alpine:latest 

RUN apk --no-cache add ca-certificates

WORKDIR /app/

COPY --from=build_stage ./app ./

CMD [ "./bin/app","-p","8080" ]
