FROM gcr.io/distroless/static

COPY scaleway-ddns /
ENTRYPOINT ["/scaleway-ddns"]
