#base image
FROM golang:1.20
#working directory
WORKDIR /demoDocker
#installing dependencies
COPY go.mod ./
RUN go mod download
#copying all files 
COPY . .
#starting app
CMD [ "go","run","main.go" ]
#exposing server port
EXPOSE 8080