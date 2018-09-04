# Standard library
from abc import ABCMeta, abstractmethod

# Internal modules
from app import db
from app.models import Stock


class StockRepo(metaclass=ABCMeta):
    """Interface for storage and retrieval of Stock entinties."""

    @abstractmethod
    def get_all(self, include_inactive=False):
        """Gets stored Stocks either only active of all.

        :return: List of all stored stocks.
        """
        pass

    @abstractmethod
    def save(self, stock):
        """Stores a new stock.

        :param stock: Stock to store.
        """
        pass


class SQLStockRepo(StockRepo):
    """StockRepo implemented against a SQL database."""

    def get_all(self, include_inactive=False):
        if include_inactive:
            return Stock.query.all()
        return Stock.query.filter(Stock.is_active == True).app()


    def save(self, stock):
        db.session.add(stock)
        db.session.commit()
