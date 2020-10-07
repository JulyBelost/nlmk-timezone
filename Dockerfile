FROM alpine:latest
WORKDIR /nlmk
RUN apk add --no-cache tzdata
COPY time /nlmk

EXPOSE 8080
CMD ["./time"]