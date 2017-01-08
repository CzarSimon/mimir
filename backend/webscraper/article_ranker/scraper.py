from datetime import datetime
from newspaper import Article
from sys import platform

import re


def fetch_page_content(url):
    article = _retrive_content(url)
    content = _format_content(article)
    return content


def _retrive_content(url):
    article = Article(url, MAX_SUMMARY_SENT=3)
    article.download()
    return _parse_text(article)


def _parse_text(raw_article):
    article = raw_article  # _handle_plattform(raw_article)
    article.parse()
    article.nlp()
    return article


def _handle_plattform(article):
    if platform == "darwin":
        regex = re.compile('(class="Emoji.*?)alt=".*?"', r'\g<1> alt=""')
        article.html = re.sub(regex, article.html)
    return article


def _format_content(article):
    return dict(
        title=article.title,
        text=article.text,
        timestamp=_timestamp(article.publish_date),
        summary=article.summary,
        img_url=article.top_image,
        keywords=article.keywords,
    )


def _timestamp(candidate):
    stamp = candidate if (candidate is not None) else datetime.utcnow()
    return stamp.strftime("%Y-%m-%d")
