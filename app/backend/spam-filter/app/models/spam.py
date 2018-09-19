# Standard library
from datetime import datetime
from enum import Enum


# Internal modules
from app import db
from app.service.training_service import format_text


class Label(Enum):
    SPAM = 'SPAM'
    NON_SPAM = 'NON-SPAM'


class SpamLabel(db.Model):
    label = db.Column(db.String(10), primary_key=True)
    created_at = db.Column(db.DateTime, default=datetime.utcnow, nullable=False)

    def __repr__(self):
        return f'SpamLabel(label={self.label} created_at={self.created_at})'


class TrainingData(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    text = db.Column(db.String(500))
    label = db.Column(db.String(50), db.ForeignKey('spam_label.label'))
    created_at = db.Column(db.DateTime, default=datetime.utcnow, nullable=False)

    def __repr__(self):
        return 'TrainingData(id={} text={} label={} created_at={})'.\
            format(self.id, self.text, self.label, self.created_at)


class SpamCandidate(object):
    def __init__(self, text, label=None):
        self.text = format_text(text)
        self.label = label

    def to_dict(self):
        return {
            'text': self.text,
            'label': self.label
        }

    @staticmethod
    def from_dict(raw_dict):
        return SpamCandidate(
            text=raw_dict['text'],
            label=raw_dict['label'] if 'label' in raw_dict else None)
