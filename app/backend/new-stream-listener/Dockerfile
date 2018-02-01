FROM maven:3.5.2-jdk-8-alpine

WORKDIR /usr/src/stream-listener
COPY . .
RUN mvn clean package -q

CMD ["java", "-jar", "target/streamlistener-0.1-jar-with-dependencies.jar"]
