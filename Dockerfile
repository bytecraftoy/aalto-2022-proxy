FROM docker.io/eclipse-temurin:17-jammy

ENV SBT_VERSION 1.7.2

RUN apt-get update && \
    apt-get install --no-install-recommends -y unzip=6.0-26ubuntu3.1 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN curl -L -o sbt-$SBT_VERSION.zip https://github.com/sbt/sbt/releases/download/v1.7.2/sbt-$SBT_VERSION.zip && \
    unzip sbt-$SBT_VERSION.zip -d ops

WORKDIR /apikeyproxy

COPY . /apikeyproxy

EXPOSE 8080

CMD ["/ops/sbt/bin/sbt", "jetty:start", "jetty:join"]
