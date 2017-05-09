from datetime import datetime
from newspaper import Article, ArticleException
from sys import platform, stderr

import re


def fetch_page_content(url):
    article, success = _retrive_content(url)
    content = _format_content(article) if success else None
    return content, success


def _retrive_content(url):
    article = Article(url)
    success = False
    try:
        article.download()
        article.parse()
        article.nlp()
        success = True
    except ArticleException as e:
        sterr.write(e)
    finally:
        return article, success # _parse_text(article)


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
