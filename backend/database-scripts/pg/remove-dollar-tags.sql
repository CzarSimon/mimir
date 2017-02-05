ALTER TABLE stocktweets DROP CONSTRAINT stocktweets_ticker_fkey;
ALTER TABLE tickeraliases DROP CONSTRAINT tickeraliases_ticker_fkey;

UPDATE stocks SET ticker=SUBSTRING(ticker, 2, LENGTH(ticker));
UPDATE tickeraliases SET ticker=SUBSTRING(ticker, 2, LENGTH(ticker));
UPDATE tickeraliases SET alias=SUBSTRING(alias, 2, LENGTH(alias));
UPDATE stocktweets SET ticker=SUBSTRING(ticker, 2, LENGTH(ticker));

ALTER TABLE stocktweets ADD FOREIGN KEY (ticker) REFERENCES stocks (ticker);
ALTER TABLE tickeraliases ADD FOREIGN KEY (ticker) REFERENCES stocks (ticker);
