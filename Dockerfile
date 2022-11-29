#
# BUILD STAGE
#
FROM docker.io/eclipse-temurin:17-jammy as build

RUN apt-get update && \
    apt-get install --no-install-recommends -y unzip=6.0-26ubuntu3.1 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENV SBT_VERSION 1.7.2
RUN curl -L -o sbt-$SBT_VERSION.zip https://github.com/sbt/sbt/releases/download/v${SBT_VERSION}/sbt-${SBT_VERSION}.zip && \
    unzip sbt-${SBT_VERSION}.zip -d /ops

ENV WORKDIR /apikeyproxy
WORKDIR ${WORKDIR}

COPY . ${WORKDIR}

RUN /ops/sbt/bin/sbt clean assembly

#
# DEPLOY STAGE
#
FROM docker.io/eclipse-temurin:17-jre-jammy

USER proxy:proxy

ENV WORKDIR /apikeyproxy
WORKDIR ${WORKDIR}

COPY --from=build --chown=proxy:proxy ${WORKDIR}/apikeyproxy.jar ${WORKDIR}

EXPOSE 8080

CMD ["java", "-jar", "apikeyproxy.jar"]
