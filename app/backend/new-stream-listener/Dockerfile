FROM maven:3.5.2-jdk-8-alpine

RUN mkdir -p /usr/src/stream-listener
COPY . /usr/src/stream-listener
WORKDIR /usr/src/stream-listener

RUN mvn clean package

CMD ["java", "-jar", "target/streamlistener-0.1-jar-with-dependencies.jar"]