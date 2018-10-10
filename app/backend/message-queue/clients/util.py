# Standard library
import os

# 3rd party modules
import pika


# Types
MQChannel = pika.adapters.blocking_connection.BlockingChannel
MQConnection = pika.BlockingConnection
MQDeliver = pika.spec.Basic.Deliver
MQProperties = pika.amqp_object.Properties


def get_queue_uri() -> str:
    user: str = os.environ['RABBITMQ_USER']
    password: str = os.environ['RABBITMQ_PASSWORD']
    return f'amqp://{user}:{password}@localhost:5672/'


def get_queue_connection(uri: str) -> MQConnection:
    return pika.BlockingConnection(pika.URLParameters(uri))


def get_channel(name: str, conn: MQConnection) -> MQChannel:
    channel = conn.channel()
    return channel
