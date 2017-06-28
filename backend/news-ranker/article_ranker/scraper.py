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
        article.build()
        success = True
    except ArticleException as e:
        sterr.write(e)
    finally:
        return article, success


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
        #img_url=article.top_image,
        keywords=article.keywords,
    )

def check_if_scraped(stored_article):
    is_scraped = stored_article["isScraped"]
    content = _stored_article_to_content(stored_article) if is_scraped else {}
    return content, is_scraped


def _stored_article_to_content(stored_article):
    return dict(
        title=stored_article["title"],
        text=stored_article["body"],
        timestamp=stored_article["dateInserted"],
        summary=stored_article["summary"],
        keywords=stored_article["keywords"]
    )


def _timestamp(candidate):
    now = datetime.utcnow()
    stamp = candidate if (candidate is not None) else now
    time_component = now.strftime("T%H:%M:%SZ")
    return stamp.strftime("%Y-%m-%d") + time_component
