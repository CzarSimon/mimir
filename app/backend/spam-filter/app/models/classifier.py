# Standard library
from datetime import datetime
from enum import Enum

# Internal modules
from app import db

class ModelType(Enum):
    SVM = 'SVM'
    NAIVE_BAYES = 'NAIVE-BAYES'


class Classifier(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    type = db.Column(db.String(50))
    training_samples = db.Column(db.Integer, nullable=False)
    test_samples = db.Column(db.Integer, nullable=False)
    accuracy = db.Column(db.Float, nullable=False)
    model_hash = db.Column(db.String(64))
    created_at = db.Column(db.DateTime, default=datetime.utcnow, nullable=False)

    def __repr__(self):
        return ('Classifier(id={} type={} training_samples={} '
                'test_samples={} accuracy={} model_hash={} '
                'created_at={})').format(
                    self.id, self.type, self.training_samples,
                    self.test_samples, self.accuracy, self.model_hash,
                    self.created_at)
