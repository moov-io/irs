FROM golang:1.18 as builder
RUN apt-get update && apt-get install -y pdftk make gcc g++ ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /src
COPY . .
COPY ./configs/config.default.yml /configs/config.default.yml
RUN make build

FROM debian:stable AS runtime
LABEL maintainer="Moov <support@moov.io>"

RUN apt-get update && apt-get install -y curl

COPY --from=builder /src/bin/* /app/
COPY --from=builder /configs/config.default.yml /configs/config.default.yml

ENV HTTP_PORT=8208
ENV HEALTH_PORT=8209

EXPOSE ${HTTP_PORT}/tcp
EXPOSE ${HEALTH_PORT}/tcp

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
	CMD curl -f http://localhost:${HEALTH_PORT}/live || exit 1

VOLUME [ "/data", "/configs" ]

ENTRYPOINT ["/app/irs"]
CMD ["web"]
