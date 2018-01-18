# lib/go/schema
Contains common data models used throughout mimir

## news
Data models related to news.
1. Article - Contains article info.
2. Author - Contains info about an article refererer.
3. RankObject - Contains info to scrape and rank an article.
4. Subject - Subject which to look for in an article.

## spam
Data models related to spam classification.
1. Candidate - Potential spam candidate and label.

## stock
Data models related stocks, identifiers and prices.
1. Stock - Contains information about a stock and issuing company.
2. Price - Holds price information about a stock.
3. Tickers - List tickers that can be queried and inserted into a sql database.

## user
Data models related to an application user
1. User - Holds user information
