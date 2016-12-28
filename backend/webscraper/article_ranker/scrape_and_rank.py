from hashlib import md5
import json
from ranker import calc_subject_scores as score
import scraper
import sys
from sys import argv


def main():
    # Variable declaration
    params = json.loads(argv[1])
    article_info = params["article_info"]
    ref_score = params["reference_score"]
    url = article_info["url"]

    # Information retrival and ranking
    content = scraper.fetch_page_content(url)
    subject_score = score(article_info["subjects"], content["text"])
    compound_score = {}
    for ticker, sub_score in subject_score.items():
        compound_score[str(ticker)] = sub_score + ref_score

    # Response construction
    response = dict(
        id=md5(url.encode('utf-8')).hexdigest(),
        url=url,
        title=content["title"],
        summary=content["summary"],
        text=content["text"],
        keywords=content["keywords"],
        subject_score=subject_score,
        reference_score=ref_score,
        compound_score=compound_score,
        timestamp=content["timestamp"],
        twitter_references=[article_info["author"]["id"]]
    )
    print(json.dumps(response))


if __name__ == "__main__":
    main()
    sys.exit(0)
