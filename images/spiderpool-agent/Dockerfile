# Copyright 2022 Authors of spidernet-io
# SPDX-License-Identifier: Apache-2.0

ARG BASE_IMAGE=ghcr.io/spidernet-io/spiderpool/spiderpool-base:1f8330482d25b58d2ae26bc6252e20384bac92ad
ARG GOLANG_IMAGE=docker.io/library/golang:1.20@sha256:f69d47fedd3b2ebd23bcf473c0b78522ebbc1823f06b7d47f45f04a30bdc901d

#======= build bin ==========
FROM --platform=${BUILDPLATFORM} ${GOLANG_IMAGE} as builder

ARG TARGETOS
ARG TARGETARCH
ARG RACE
ARG NOSTRIP
ARG NOOPT
ARG QUIET_MAKE

COPY . /src
WORKDIR /src/cmd/spiderpool-agent
RUN  make GOARCH=${TARGETARCH}   \
        RACE=${RACE} NOSTRIP=${NOSTRIP} NOOPT=${NOOPT} QUIET_MAKE=${QUIET_MAKE} \
        DESTDIR_BIN=/tmp/install/${TARGETOS}/${TARGETARCH}/bin \
        all install

WORKDIR /src/cmd/spiderpool
RUN  make GOARCH=${TARGETARCH}   \
        RACE=${RACE} NOSTRIP=${NOSTRIP} NOOPT=${NOOPT} QUIET_MAKE=${QUIET_MAKE} \
        DESTDIR_BIN=/tmp/install/${TARGETOS}/${TARGETARCH}/bin \
        all install

WORKDIR /src/cmd/coordinator
RUN make GOARCH=${TARGETARCH}   \
        RACE=${RACE} NOSTRIP=${NOSTRIP} NOOPT=${NOOPT} QUIET_MAKE=${QUIET_MAKE} \
        DESTDIR_BIN=/tmp/install/${TARGETOS}/${TARGETARCH}/bin \
        all install

#====== release image =======

FROM ${BASE_IMAGE}

LABEL maintainer="maintainer@spidernet-io"

# TARGETOS is an automatic platform ARG enabled by Docker BuildKit.
ARG TARGETOS
# TARGETARCH is an automatic platform ARG enabled by Docker BuildKit.
ARG TARGETARCH

ARG GIT_COMMIT_VERSION
ENV GIT_COMMIT_VERSION=${GIT_COMMIT_VERSION}
ARG GIT_COMMIT_TIME
ENV GIT_COMMIT_TIME=${GIT_COMMIT_TIME}
ARG VERSION
ENV VERSION=${VERSION}

COPY --from=builder /tmp/install/${TARGETOS}/${TARGETARCH}/bin/*   /usr/bin/

CMD ["/usr/bin/spiderpool-agent daemon"]
