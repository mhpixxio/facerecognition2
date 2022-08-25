# syntax=docker/dockerfile:1

#set base image
FROM golang:1.18.3-bullseye

#set workspace
WORKDIR /app

#update the repository sources list
RUN apt-get update -qq

#add other stuff
RUN apt-get install --no-install-recommends libvips42 -y
RUN apt-get install libdlib-dev libblas-dev libatlas-base-dev liblapack-dev libjpeg62-turbo-dev -y
RUN apt-get install imagemagick imagemagick-doc -y

#add the model files
COPY /models/*.dat ./models/

#add the go files
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./

#run the go code
RUN go build -o /docker-face-recognition

#expose the port
EXPOSE 8080

CMD ["/docker-face-recognition"]

