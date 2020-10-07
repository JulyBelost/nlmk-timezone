FROM alpine:latest
WORKDIR /nlmk
COPY time /nlmk

EXPOSE 8080
CMD ["./time"]