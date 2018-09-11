# Standard library
from datetime import datetime
from enum import Enum


# Internal modules
from app import db


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


class Classifier(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    training_samples = db.Column(db.Integer, nullable=False)
    test_samples = db.Column(db.Integer, nullable=False)
    accuracy = db.Column(db.Float, nullable=False)
    model_hash = db.Column(db.String(64))
    created_at = db.Column(db.DateTime, default=datetime.utcnow, nullable=False)

    def __repr__(self):
        return ('Classifier(id={} training_samples={} test_samples={} '
                'accuracy={} model_hash={} created_at={})').format(
                    self.id, self.training_samples, self.test_samples,
                    self.accuracy, self.model_hash, self.created_at)
