# build stage
FROM golang:1.22-alpine AS build

# set working directory
WORKDIR /app

# copy source code
COPY . .

# install dependencies
RUN go mod download

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/crm main.go

# final stage
FROM alpine:latest AS final
LABEL maintainer="Roihan Sori"

# set working directory
WORKDIR /app

# copy binary
COPY --from=build /app/bin/crm ./

COPY .env /app

EXPOSE 5000

ENTRYPOINT [ "./crm" ]