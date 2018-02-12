import json
from datetime import datetime
import hashlib


def new_id():
    time_str = str(datetime.utcnow())
    return hashlib.md5(time_str).hexdigest()


tweet = {
    "contributors": None,
    "coordinates": None,
    "created_at": "Tue Aug 29 06:15:45 +0000 2017",
    "entities": {
        "hashtags": [],
        "symbols": [
            {
                "indices": [
                    67,
                    72
                ],
                "text": "ORCL"
            },
            {
                "indices": [
                    67,
                    72
                ],
                "text": "TSLA"
            },
            {
                "indices": [
                    67,
                    72
                ],
                "text": "NOT"
            },
            {
                "indices": [
                    67,
                    72
                ],
                "text": "GOOGL"
            }
        ],
        "urls": [
            {
                "display_url": "ift.tt/2wOjrVh",
                "expanded_url": "http://ift.tt/2wOjrVh",
                "indices": [
                    73,
                    96
                ],
                "url": "https://t.co/xEB120bOF2"
            }
        ],
        "user_mentions": []
    },
    "favorite_count": 0,
    "favorited": False,
    "filter_level": "low",
    "geo": None,
    "id": 902414451567198208,
    "id_str": new_id(),
    "in_reply_to_screen_name": None,
    "in_reply_to_status_id": None,
    "in_reply_to_status_id_str": None,
    "in_reply_to_user_id": None,
    "in_reply_to_user_id_str": None,
    "is_quote_status": False,
    "lang": "en",
    "place": None,
    "possibly_sensitive": False,
    "quote_count": 0,
    "reply_count": 0,
    "retweet_count": 0,
    "retweeted": False,
    "source": "<a href=\"https://ifttt.com\" rel=\"nofollow\">IFTTT</a>",
    "text": "Weekly Investment Analysts\u2019 Ratings Updates for Oracle Corporation $ORCL https://t.co/xEB120bOF2",
    "timestamp_ms": "1503987345960",
    "truncated": False,
    "user": {
        "contributors_enabled": False,
        "created_at": "Fri Jan 15 17:15:13 +0000 2016",
        "default_profile": True,
        "default_profile_image": False,
        "description": None,
        "favourites_count": 0,
        "follow_request_sent": None,
        "followers_count": 154,
        "following": None,
        "friends_count": 36,
        "geo_enabled": False,
        "id": 4816109896,
        "id_str": "4816109896",
        "is_translator": False,
        "lang": "en",
        "listed_count": 16,
        "location": None,
        "name": "EMQ News",
        "notifications": None,
        "profile_background_color": "F5F8FA",
        "profile_background_image_url": "",
        "profile_background_image_url_https": "",
        "profile_background_tile": False,
        "profile_image_url": "http://pbs.twimg.com/profile_images/688048912293167104/UmUO1efq_normal.png",
        "profile_image_url_https": "https://pbs.twimg.com/profile_images/688048912293167104/UmUO1efq_normal.png",
        "profile_link_color": "1DA1F2",
        "profile_sidebar_border_color": "C0DEED",
        "profile_sidebar_fill_color": "DDEEF6",
        "profile_text_color": "333333",
        "profile_use_background_image": True,
        "protected": False,
        "screen_name": "emq_news",
        "statuses_count": 185081,
        "time_zone": None,
        "translator_type": "none",
        "url": None,
        "utc_offset": None,
        "verified": False
    }
}

raw_tweet = json.dumps(tweet)

tracking_data = {
    'tickers': set(['$CL', '$DIS', '$NKE', '$AMZN', '$PYPL', '$ZNGA', '$MON', '$AAPL', '$GS', '$MSFT', '$YELP', '$SQ', '$SNAP', '$ORCL', '$M', '$F', '$UA', '$GE', '$TWTR', '$BABA', '$JPM', '$GOOG', '$FIT', '$NVDA', '$T', '$S', '$FB', '$CSCO', '$ACN', '$BA', '$AXP', '$VZ', '$V', '$INTC', '$CRM', '$NFLX', '$GM', '$AMD', '$HPQ', '$EBAY', '$WMT', '$C', '$GPRO', '$PG', '$TSLA', '$BIDU']),
    'stock_querys': {
        'GOOG': {'ticker': 'GOOG', 'name': 'Alphabet Inc.'},
        'AXP': {'ticker': 'AXP', 'name': 'American Express Company'},
        'NFLX': {'ticker': 'NFLX', 'name': 'Netflix, Inc.'},
        'BA': {'ticker': 'BA', 'name': 'Boeing Company'},
        'TSLA': {'ticker': 'TSLA', 'name': 'Tesla Motors, Inc.'},
        'AAPL': {'ticker': 'AAPL', 'name': 'Apple Inc.'},
        'ACN': {'ticker': 'ACN', 'name': 'Accenture plc'},
        'ZNGA': {'ticker': 'ZNGA', 'name': 'Zynga Inc.'},
        'FB': {'ticker': 'FB', 'name': 'Facebook Inc.'},
        'GM': {'ticker': 'GM', 'name': 'General Motors Company'},
        'NVDA': {'ticker': 'NVDA', 'name': 'NVIDIA Corporation'},
        'GPRO': {'ticker': 'GPRO', 'name': 'GoPro, Inc.'},
        'AMZN': {'ticker': 'AMZN', 'name': 'Amazon.com, Inc.'},
        'MSFT': {'ticker': 'MSFT', 'name': 'Microsoft Corporation'},
        'DIS': {'ticker': 'DIS', 'name': 'Walt Disney Company'},
        'FIT': {'ticker': 'FIT', 'name': 'Fitbit, Inc.'},
        'TWTR': {'ticker': 'TWTR', 'name': 'Twitter, Inc.'},
        'F': {'ticker': 'F', 'name': 'Ford Motor Company'},
        'AMD': {'ticker': 'AMD', 'name': 'Advanced Micro Devices, Inc.'},
        'ORCL': {'ticker': 'ORCL', 'name': 'Oracle Corporation'},
        'PG': {'ticker': 'PG', 'name': 'Procter & Gamble Company'},
        'CRM': {'ticker': 'CRM', 'name': 'Salesforce.com Inc'},
        'BABA': {'ticker': 'BABA', 'name': 'Alibaba Group Holding Limited'},
        'C': {'ticker': 'C', 'name': 'Citigroup, Inc.'},
        'GS': {'ticker': 'GS', 'name': 'The Goldman Sachs Group, Inc.'},
        'INTC': {'ticker': 'INTC', 'name': 'Intel Corporation'},
        'CL': {'ticker': 'CL', 'name': 'Colgate-Palmolive Company'},
        'PYPL': {'ticker': 'PYPL', 'name': 'PayPal Holdings, Inc.'},
        'HPQ': {'ticker': 'HPQ', 'name': 'HP Inc.'},
        'M': {'ticker': 'M', 'name': "Macy's Inc"},
        'WMT': {'ticker': 'WMT', 'name': 'Wal-Mart Stores Inc.'},
        'GE': {'ticker': 'GE', 'name': 'General Electric Company'},
        'T': {'ticker': 'T', 'name': 'AT&T, Inc.'},
        'V': {'ticker': 'V', 'name': 'Visa Inc.'},
        'VZ': {'ticker': 'VZ', 'name': 'Verizon Communications Inc.'},
        'JPM': {'ticker': 'JPM', 'name': 'JP Morgan Chase & Co.'},
        'NKE': {'ticker': 'NKE', 'name': 'Nike, Inc.'},
        'YELP': {'ticker': 'YELP', 'name': 'Yelp Inc.'},
        'BIDU': {'ticker': 'BIDU', 'name': 'Baidu, Inc.'},
        'CSCO': {'ticker': 'CSCO', 'name': 'Cisco Systems, Inc.'},
        'S': {'ticker': 'S', 'name': 'Sprint Corporation'},
        'EBAY': {'ticker': 'EBAY', 'name': 'eBay Inc.'},
        'MON': {'ticker': 'MON', 'name': 'Monsanto Company'},
        'SNAP': {'ticker': 'SNAP', 'name': 'Snap Inc.'},
        'UA': {'ticker': 'UA', 'name': 'Under Armour, Inc.'},
        'SQ': {'ticker': 'SQ', 'name': 'Square, Inc.'}
    },
    'aliases': {
        'GOOGL': 'GOOG'
    }
}
