import scraper
import json
import gc
import sys
from sys import argv
from hashlib import md5
from datetime import datetime
from ranker import calc_subject_scores

def main():
    # Variable declaration
    params = json.loads(argv[1])
    article_info = params["article_info"]
    ref_score = params["reference_score"]
    url = article_info["url"]

    # Information retrival and ranking
    content = scraper.fetch_page_content(url)
    subject_score = calc_subject_scores(article_info["subjects"], content["text"])
    compound_score = {}
    for ticker, sub_score in subject_score.iteritems():
        compound_score[str(ticker)] = sub_score + ref_score

    # Response construction
    response = {
        "id": md5(url).hexdigest(),
        "url": url,
        "title": content["title"],
        "subject_score": subject_score,
        "reference_score": ref_score,
        "compound_score": compound_score,
        "timestamp": datetime.utcnow().strftime("%Y-%m-%d"),
        "twitter_references": [article_info["author"]["id"]]
    }
    print json.dumps(response)

if __name__ == "__main__":
    main()
    gc.collect()
    sys.exit(0)
