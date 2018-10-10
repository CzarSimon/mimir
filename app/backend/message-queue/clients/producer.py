# Standard library
import json
import os
from dataclasses import dataclass, asdict

# Internal packages
import util
from model import Message

EXCANGE_NAME: str = 'x-news'
CHANNEL_NAME: str = 'q-rank-objects'
conn: util.MQConnection = util.get_queue_connection(util.get_queue_uri())
channel: util.MQChannel = conn.channel()


def send_message(msg: Message) -> None:
    print(f'Sent message: {msg.tojson()}')
    channel.basic_publish(exchange=EXCANGE_NAME,
                          routing_key=CHANNEL_NAME,
                          body=msg.tojson())


def main() -> None:
    for i in range(100):
        send_message(Message(body=f'Message nr: {i}', number=i))
    conn.close()


if __name__ == '__main__':
    main()
