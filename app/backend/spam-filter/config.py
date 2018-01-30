# -*- coding: utf-8 -*-
import os
import sys



def _get_cashtag_threshold():
    threshold_str = os.getenv("CASHTAG_THRESHOLD", "0.8")
    try:
        return float(threshold_str)
    except:
        print "{}: could not parse value as float".format(threshold_str)
        sys.exit(1)


cashtag_threshold = _get_cashtag_threshold()


database = {
    "NAME": os.getenv("DB_NAME", "mimirprod"),
    "USER": os.getenv("DB_USER","simon"),
    "HOST": os.getenv("DB_HOST", "localhost"),
    "PASSWORD": os.getenv("DB_PASSWORD", "pwd"),
    "PORT": os.getenv("DB_PORT", "5432")
}


categories = {
    "SPAM": 'SPAM',
    "NON-SPAM": 'NON-SPAM'
}
