from hashlib import md5
import json
from ranker import calc_subject_scores, calc_compound_score
import scraper
import req
from sys import argv, stderr


def create_response(content, subject_score, params):
    url = params["articleInfo"]["url"]
    newArticle = dict(
        id=md5(url.encode('utf-8')).hexdigest(),
        url=url,
        title=content["title"],
        summary=content["summary"],
        text=content["text"],
        keywords=content["keywords"],
        subject_score=subject_score,
        reference_score=params["referenceScore"],
        compound_score=calc_compound_score(params["referenceScore"], subject_score),
        timestamp=content["timestamp"],
        twitter_references=[params["articleInfo"]["author"]["id"]]
    )
    return dict(
        newArticle=newArticle,
        storedArticle=params["storedArticle"]
    )


def send_response(response):
    json_response = json.dumps(response)
    req.send_response(json_response)
    #print(json_response)


def main():
    # Variable declaration
    params = json.loads(argv[1])
    article_info = params["articleInfo"]

    # Information retrival and ranking
    content, success = scraper.fetch_page_content(article_info["url"])
    if success:
        # Score article content score against subjects
        subject_score = calc_subject_scores(article_info["subjects"], content["text"])
        response = create_response(content, subject_score, params)
        send_response(response)
    else:
        print("Article scraping not successful")
    print("done")

if __name__ == "__main__":
    main()
