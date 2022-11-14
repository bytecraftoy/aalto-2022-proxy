FROM openjdk:8

ENV SBT_VERSION 1.7.2

RUN curl -L -o sbt-$SBT_VERSION.zip https://github.com/sbt/sbt/releases/download/v1.7.2/sbt-$SBT_VERSION.zip
RUN unzip sbt-$SBT_VERSION.zip -d ops

WORKDIR /apikeyproxy

COPY . /apikeyproxy

EXPOSE 8080

CMD /ops/sbt/bin/sbt jetty:start jetty:join