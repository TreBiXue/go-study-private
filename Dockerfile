FROM golang:1.19.1-bullseye AS build

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

RUN apt-get -qq update && apt-get -yqq install upx

WORKDIR /src

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN go build \
  -a \
  -trimpath \
  -ldflags "-s -w -extldflags '-static'" \
  -o /bin/app ./cmd/

RUN strip /bin/app
RUN upx -q -9 /bin/app


FROM gcr.io/distroless/static

COPY --from=build /bin/app /bin/app

ENTRYPOINT ["/bin/app"]