FROM ubuntu/apache2:latest

RUN apt-get update && apt-get install -y ca-certificates openssl

ARG cert_location=/usr/local/share/ca-certificates

RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
RUN update-ca-certificates
RUN mkdir forum
WORKDIR /forum
COPY forum/ ./
RUN apt update
RUN apt -y install golang-go
CMD ["go", "run","main/main.go"]
EXPOSE 3333 