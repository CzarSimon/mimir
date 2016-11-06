APP_SERVER = {
    "ADDRESS": "http://139.59.159.73",
    "routes": {
        "STOCKLIST": "/stockList"
    }
}

NEWS_SERVER = {
    "ADDRESS": "http://139.59.214.5:5000",
    "routes": {
        "RANK": "/rankArticle"
    }
}

timing = {
    "TIMEOUT": 0.5
}

# perhaps include dlvr.it here?
forbidden_domains = ["owler.us", "owler.com", "stocktwits.com", "investorshangout.com", "1broker.com"]
