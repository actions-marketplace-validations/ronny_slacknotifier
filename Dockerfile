# Based on https://www.sethvargo.com/writing-github-actions-in-go/
#############################################################################
FROM golang:1.16 AS builder

RUN apt-get update && apt-get -y install upx

ENV CGO_ENABLED=0

WORKDIR /src

# These layers shouldn't change if there are no dependency changes
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make binaries

# Strip any symbols - this is not a library
RUN strip bin/*

# Compress the compiled action
RUN upx -q -9 bin/*

#############################################################################
FROM scratch

# Copy over SSL certificates from the first step - this is required
# if our code makes any outbound SSL connections because it contains
# the root CA bundle.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /src/bin/ /bin/

CMD ["/bin/action-slack-notify"]
