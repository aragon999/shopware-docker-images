FROM --platform=$BUILDPLATFORM golang:1.27.3-alpine AS builder

ARG TARGETOS
ARG TARGETARCH
ENV GOTOOLCHAIN=local
WORKDIR /src

# Build upstream supervisord from a pinned, verified source archive.
ADD --checksum=sha256:2641a8772113878d3ff526151368a9ba7ac979b725ef107920b8b5cde44e5cf2 \
    https://github.com/ochinchina/supervisord/archive/bc67b81f8475434fdf31e20ececb273feb3252cf.tar.gz /tmp/supervisord.tar.gz
RUN tar -xzf /tmp/supervisord.tar.gz --strip-components=1 && \
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -mod=readonly -buildvcs=false -trimpath -ldflags="-s -w" -o /out/supervisord .

FROM scratch

LABEL org.opencontainers.image.source=https://github.com/aragon999/shopware-docker-images

COPY --from=builder /out/supervisord /usr/local/bin/supervisord
ENTRYPOINT ["/usr/local/bin/supervisord"]
