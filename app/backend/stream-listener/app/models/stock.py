# Standard library
from datetime import datetime

# 3rd party modules
import sqlalchemy as sa

# Internal modules
from app import db


class Stock(db.Model):
    __tablename__ = 'stock'

    symbol = sa.Column(sa.String(20), primary_key=True)
    name = sa.Column(sa.String(100))
    is_active = sa.Column(sa.Boolean)
    total_count =sa. Column(sa.Integer)
    updated_at = sa.Column(sa.DateTime)

    def __init__(self, symbol, name, is_active, total_count, updated_at=None):
        self.symbol = symbol
        self.name = name
        self.is_active = is_active
        self.total_count = total_count
        self.updated_at = updated_at or datetime.utcnow()

    def __repr__(self):
        return ('Stock(symbol={} name={} is_active={} '
                'total_count={} updated_at={})'.format(
                self.symbol, self.name, self.is_active,
                self.total_count, self.updated_at))
