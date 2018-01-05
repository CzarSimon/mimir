from hashlib import sha256
import req
import date


schema = {
    "urlHash": "string",
    "title": "string",
    "ticker": "string",
    "date": "string",
    "score": {
        "subjectScore": "float",
        "referenceScore": "float"
    }
}


def new_article(title, ticker, url, score):
    article = dict(
        urlHash=sha256(url).hexdigest(),
        title=title,
        ticker=ticker,
        date=date.utc_str(),
        score=new_score(score)
    )
    print article["title"], article["urlHash"]
    return article


def new_score(score):
    return dict(
        subjectScore=score,
        referenceScore=(2 * score)
    )


def test_articles():
    return [
        new_article("title 1", "a", "1", 0.1),
        new_article("Title 1", "A", "2", 0.3),
        new_article("title 2", "A", "3", 0.2),
        new_article("title 2", "B", "4", 0.1),
        new_article("title 1", "b", "5", 0.05),
        new_article("title 1", "b", "6", 0.7),
        new_article("title 2", "A", "7", 0.8),
        new_article("title 1", "A", "8", 0.4),
        new_article("TITLE 2", "b", "9", 0.15)
    ]


def main():
    articles = test_articles()
    url = "http://localhost:6000/api/cluster-article"
    for article in articles:
        res = req.post_data(article, url)
        print res


if __name__ == '__main__':
    main()
