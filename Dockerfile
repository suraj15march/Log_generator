FROM golang:alpine
RUN mkdir /app
RUN mkdir /var/log/myapp
RUN chmod -R a+rwx /var/log/myapp
COPY . /app
WORKDIR /app
RUN go build -o main . 
CMD ["/app/main"]