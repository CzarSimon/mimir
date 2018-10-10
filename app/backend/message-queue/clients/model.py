# Standard library
import json
from dataclasses import dataclass, asdict


@dataclass(frozen=True)
class Message:
    body: str
    number: int

    def tojson(self) -> str:
        return json.dumps(asdict(self))


def parse_message(raw: bytes) -> Message:
    serialized_msg = json.loads(raw)
    return Message(
        body=serialized_msg['body'],
        number=serialized_msg['number'])
