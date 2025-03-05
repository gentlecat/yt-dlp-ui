FROM golang:alpine AS build

RUN apk update && apk add git make python3 zip ffmpeg libffi-dev

WORKDIR /build
RUN git clone --depth 1 https://github.com/yt-dlp/yt-dlp
WORKDIR /build/yt-dlp/
RUN make yt-dlp && cp yt-dlp /usr/local/bin/

WORKDIR /build/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /usr/local/bin/ ./...


FROM alpine:latest AS runtime

LABEL org.opencontainers.image.source=https://github.com/gentlecat/yt-dlp-ui
LABEL org.opencontainers.image.description="UI for yt-dlp"
LABEL org.opencontainers.image.licenses=MIT

EXPOSE 8080

COPY --from=build /usr/local/bin/yt-dlp-ui /usr/local/bin/

CMD ["/usr/local/bin/yt-dlp-ui"]
