FROM gcr.io/distroless/base

COPY scaleway-ddns /
ENTRYPOINT ["/scaleway-ddns"]
