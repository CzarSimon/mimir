# Standard library
import json
import os
from dataclasses import dataclass, asdict

# Internal packages
import util
from model import Message, parse_message


CHANNEL_NAME: str = 'q-rank-objects'
conn: util.MQConnection = util.get_queue_connection(util.get_queue_uri())
channel: util.MQChannel = conn.channel()


def handle_message(
        channel: util.MQChannel, method: util.MQDeliver,
        properties: util.MQProperties, body: bytes) -> None:
    print(json.loads(body))
    send_ack(channel, method)


def send_ack(channel: util.MQChannel, method: util.MQDeliver) -> None:
    channel.basic_ack(delivery_tag=method.delivery_tag)



def main() -> None:
    channel.basic_qos(prefetch_count=1)
    channel.basic_consume(handle_message, queue=CHANNEL_NAME)
    channel.start_consuming()
    conn.close()


if __name__ == '__main__':
    main()
