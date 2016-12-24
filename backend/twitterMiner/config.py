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

database = {
    "NAME": "mimirprod",
    "USER": "simon",
    "HOST": "localhost",
    "PASSWORD": "56error78",
    "PORT": ""
}

twitter_credentials = {
    consumerKey = "Lt9RkL58dttJ844S5z4438VVX"
    consumerSecret = "oD1q9GFJPt5B7ZnhbnLHI9Rc48srYVhMssUGaQ8ujCpIvzNvzP"
    accessToken = "2415072665-XLxE75HDMBYCHlroyZLoJgigNi4yweHIR3pLnqW"
    accessSecret = "aBT77jhPu7edRploBEtKQB6oqZIEMbbBhRbz2qjKIlMna"
}

# perhaps include dlvr.it here?
# Should be chaged to set for faster lookup.
forbidden_domains = set(["owler.us", "owler.com", "stocktwits.com", "investorshangout.com", "1broker.com"])
