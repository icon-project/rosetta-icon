###################################
# Compile golang
FROM ubuntu:20.04 as golang-builder

RUN mkdir -p /app /work
WORKDIR /work

RUN apt-get update \
  && apt-get install -y curl make gcc g++ git \
  && rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.17.8
ENV GOLANG_DOWNLOAD_SHA256 980e65a863377e69fd9b67df9d8395fd8e93858e7a24c9f55803421e453f4f99
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
  && echo "$GOLANG_DOWNLOAD_SHA256 golang.tar.gz" | sha256sum -c - \
  && tar -C /usr/local -xzf golang.tar.gz \
  && rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

######################################
# Compile rocksdb
FROM golang-builder as rocksdb-builder

RUN apt-get update \
  && apt-get install -y zlib1g-dev libbz2-dev libsnappy-dev \
  && rm -rf /var/lib/apt/lists/*

ENV ROCKSDB_VERSION v6.22.1

RUN git clone https://github.com/facebook/rocksdb.git \
  && cd rocksdb \
  && git checkout "$ROCKSDB_VERSION" \
  && PORTABLE=1 make shared_lib

RUN cd rocksdb \
  && mkdir lib \
  && cp -P librocksdb.so* lib/ \
  && cp -P lib/* /usr/lib/ \
  && cp -rP include/* /usr/include/

######################################
# Compile goloop
FROM rocksdb-builder as goloop-builder

ENV GOLOOP_VERSION v1.2.13

RUN git clone https://github.com/icon-project/goloop.git \
  && cd goloop \
  && git checkout "$GOLOOP_VERSION" \
  && make goloop

RUN mv goloop/bin/goloop /app/goloop

###################################
# Compile python and java executors
FROM goloop-builder as exec-builder

RUN apt-get update \
  && DEBIAN_FRONTEND=noninteractive apt-get install -y unzip python3.8-venv openjdk-11-jdk-headless \
  && rm -rf /var/lib/apt/lists/*

SHELL ["/bin/bash", "-c"]
RUN cd goloop \
  && python3 -m venv /app/venv \
  && source /app/venv/bin/activate \
  && pip install --upgrade pip \
  && pip install -r iconee/requirements.txt \
  && make iconexec \
  && pip install build/iconee/dist/iconee-*.whl

ENV JAVA_HOME /usr/lib/jvm/java-11-openjdk-amd64
ENV JAVAEE_VERSION 0.9.2
RUN cd goloop/javaee \
  && make javaexec \
  && unzip -q app/execman/build/distributions/execman-${JAVAEE_VERSION}.zip -d /app/ \
  && mv /app/execman-${JAVAEE_VERSION} /app/execman

######################################
# Compile rosetta-icon
FROM golang-builder as rosetta-builder

COPY . src
RUN cd src \
  && make build

RUN mv src/bin/rosetta-icon /app/rosetta-icon \
  && mv src/goloop-conf /app/ \
  && rm -rf src

###################
# Build Final Image
FROM ubuntu:20.04

RUN apt-get update \
  && apt-get install -y python3.8-venv openjdk-11-jre-headless ca-certificates libsnappy1v5 jq \
  && update-ca-certificates \
  && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /app/bin \
  && chown -R nobody:nogroup /app \
  && mkdir -p /data \
  && chown -R nobody:nogroup /data

WORKDIR /app

# Copy necessary files
COPY --from=rocksdb-builder /work/rocksdb/lib /usr/lib/
COPY --from=goloop-builder /app/goloop /app/bin/goloop
COPY --from=exec-builder /app/venv /app/venv
COPY --from=exec-builder /app/execman /app/execman
COPY --from=rosetta-builder /app/rosetta-icon /app/bin/rosetta-icon
COPY --from=rosetta-builder /app/goloop-conf /app/goloop-conf
RUN mv /app/goloop-conf/start.sh /app/bin/ \
  && chmod +x /app/bin/start.sh

# Set executable path
ENV PATH $PATH:/app/bin
ENV JAVA_HOME /usr/lib/jvm/java-11-openjdk-amd64

# Env for goloop entrypoint
ENV GOLOOP_NODE_DIR=/data/data
ENV GOLOOP_CONFIG=/data/config/server.json
ENV GOLOOP_KEY_STORE=/data/config/keystore.json
ENV GOLOOP_KEY_SECRET=/data/config/keysecret
ENV GOLOOP_P2P_LISTEN=":7080"
ENV GOLOOP_RPC_ADDR=":9080"
ENV GOLOOP_ENGINES="python,java"
ENV GOLOOP_CONSOLE_LEVEL=info
ENV GOLOOP_LOG_WRITER_FILENAME=/data/logs/goloop.log
ENV GOLOOP_LOG_WRITER_MAXAGE=14
ENV GOLOOP_LOG_WRITER_MAXSIZE=1024
ENV GOLOOP_LOG_WRITER_COMPRESS=true
ENV JAVAEE_BIN /app/execman/bin/execman
ENV PYEE_VERIFY_PACKAGE="true"

# entrypoint
RUN { \
        echo '#!/bin/bash'; \
        echo 'set -e'; \
        echo 'if [ "$GOLOOP_CONFIG" != "" ] && [ ! -f "$GOLOOP_CONFIG" ]; then'; \
        echo '  UNSET="GOLOOP_CONFIG"' ; \
        echo '  CMD="goloop server save $GOLOOP_CONFIG"'; \
        echo '  if [ "$GOLOOP_KEY_SECRET" != "" ] && [ ! -f "$GOLOOP_KEY_SECRET" ]; then'; \
        echo '    mkdir -p $(dirname $GOLOOP_KEY_SECRET)'; \
        echo '    echo -n $(date|md5sum|head -c16) > $GOLOOP_KEY_SECRET' ; \
        echo '  fi'; \
        echo '  if [ "$GOLOOP_KEY_STORE" != "" ] && [ ! -f "$GOLOOP_KEY_STORE" ]; then'; \
        echo '    UNSET="$UNSET GOLOOP_KEY_STORE"' ; \
        echo '    CMD="$CMD --save_key_store=$GOLOOP_KEY_STORE"' ; \
        echo '  fi'; \
        echo '  sh -c "unset $UNSET ; $CMD"' ; \
        echo 'fi'; \
        echo ; \
        echo 'source /app/venv/bin/activate'; \
        echo 'if [ "${GOLOOP_LOG_WRITER_FILENAME}" != "" ]; then'; \
        echo '  GOLOOP_LOGDIR=$(dirname ${GOLOOP_LOG_WRITER_FILENAME})'; \
        echo '  if [ ! -d "${GOLOOP_LOGDIR}" ]; then'; \
        echo '     mkdir -p ${GOLOOP_LOGDIR}'; \
        echo '  fi'; \
        echo 'fi'; \
        echo 'exec "$@"'; \
    } > /entrypoint \
    && chmod +x /entrypoint
ENTRYPOINT ["/entrypoint"]

CMD ["start.sh"]
