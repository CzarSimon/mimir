import requests
from bs4 import BeautifulSoup

def fetch_page_content(url):
    content = {}
    html = _fetch_html(url)
    soup = BeautifulSoup(html, "html.parser")
    content["title"] = soup.title.get_text()
    content["text"] = " ".join([content["title"], _parse_text(soup)])
    return content

def _fetch_html(url):
    try:
        res = requests.get(url)
    except:
        return
    return res.content

def _parse_text(page):
    for unwanted in page(["script", "style", "ol", "ul", "footer"]):
        unwanted.extract()

    text = page.body.get_text()
    lines = (line.strip() for line in text.splitlines())
    chunks = (phrase.strip() for line in lines for phrase in line.split("  "))
    return ' '.join(chunk for chunk in chunks if chunk)





