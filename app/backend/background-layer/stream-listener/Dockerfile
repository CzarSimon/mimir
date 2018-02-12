FROM maven:3.5.2-jdk-8-alpine

WORKDIR /usr/src/stream-listener
COPY . .
RUN mvn clean package

WORKDIR /usr/src/stream-listener/target
CMD ["java", "-jar", "streamlistener-1.0-jar-with-dependencies.jar"]
