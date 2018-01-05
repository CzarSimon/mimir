# -*- coding: utf-8 -*-
import os


def get_env(key, default=""):
    return os.environ[key] if (key in os.environ) else default


server = {
    "PORT": "1000"
}


postgres = {
    "NAME": "mimirprod",
    "USER": "simon",
    "PASSWORD": get_env("PG_PASSWORD", "56error78"),
    "HOST": get_env("PG_HOST", "localhost"),
    "PORT": get_env("PG_PORT", "5432")
}


categories = {
    "SPAM": 'SPAM',
    "NON-SPAM": 'NON-SPAM'
}

cashtag_threshold = 0.8
