FROM centurylink/ca-certs
COPY goweb /
ENTRYPOINT ["/goweb"]


