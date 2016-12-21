FROM scratch
EXPOSE 80
COPY goweb /
ENTRYPOINT ["/goweb"]

