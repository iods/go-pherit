ARG GOLANG_VER=1.15
FROM golang:$GOLANG_VER

# Default labels & meta information
LABEL \
    maintaner="Rye Miller" \
    version="0.0"

# Set Environment Vars
ENV PORT=8000

WORKDIR /app
# COPY go.mod .
# RUN go mod download

COPY . .

RUN \
    cd cmd/roman \
    && go build .

CMD ["./cmd/roman/roman"]