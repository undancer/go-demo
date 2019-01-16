FROM scratch

ADD build/app-linux-amd64 /app/app

CMD ["/app/app"]