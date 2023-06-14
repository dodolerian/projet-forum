FROM golang:latest
COPY ./forum ./forum
WORKDIR /forum
RUN cd /forum
RUN go mod init forum
RUN go mod tidy
CMD ["go", "run","main/main.go"]
EXPOSE 3333 