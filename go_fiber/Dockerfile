FROM golang:alpine

# meta
LABEL name="redesiigner-api"
LABEL maintainer="redesiigner"
LABEL author="redesiigner"

# arguments

ARG DB_USERNAME=username
ARG DB_PASSWORD=password

# copy source files
WORKDIR /src
COPY src ./

# build
RUN go install
RUN go build -v -o app

# run
EXPOSE $PORT
ENTRYPOINT [ "/src/app" ]