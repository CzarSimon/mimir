# Standard library
from threading import Thread

# 3rd party modules.
from sqlalchemy import create_engine

# Internal modules
from app.config import values, DBConfig, HealthCheckConfig
from app.service import emit_heartbeats
from .database import Database

health_config = HealthCheckConfig()
Thread(
    target=emit_heartbeats,
    args=(health_config.FILENAME, health_config.INTERVAL,)).start()

db = Database(DBConfig())


from .app import App
stream_listner = App()
