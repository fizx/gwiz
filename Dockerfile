FROM golang:latest 
RUN git --version

RUN mkdir /app 
ADD go.mod /app/
ADD go.sum /app/
WORKDIR /app/ 
RUN go mod download
ADD . /app/
RUN go build .