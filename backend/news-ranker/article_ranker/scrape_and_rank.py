import json
from ranker import calc_subject_scores, calc_compound_score
import scraper
import req
from sys import argv, stderr


def create_response(content, subject_score, params):
    newArticle = dict(
        title=content["title"],
        summary=content["summary"],
        body=content["text"],
        keywords=content["keywords"],
        subjectScore=subject_score,
        compoundScore=calc_compound_score(params["referenceScore"], subject_score),
        timestamp=content["timestamp"]
    )
    return dict(
        newArticle=newArticle,
        storedArticle=params["storedArticle"]
    )

    #id=params["storedArticle"]["urlHash"],
    #url=params["url"],
    #referenceScore=params["referenceScore"],
    #twitterReferences=[params["articleInfo"]["author"]["id"]]


def send_response(response):
    json_response = json.dumps(response)
    req.send_response(json_response)


def main():
    # Variable declaration
    params = json.loads(argv[1])
    article_info = params["articleInfo"]

    content, already_scraped = scraper.check_if_scraped(params["storedArticle"])
    if not already_scraped:
        # Information retrival and ranking
        content, success = scraper.fetch_page_content(params["url"])
    else:
        success = True

    if success:
        # Score article content score against subjects
        subject_score = calc_subject_scores(article_info["subjects"], content["text"], params["storedArticle"]["urlHash"])
        response = create_response(content, subject_score, params)
        send_response(response)
    else:
        print("Article scraping not successful")
    print("done")

if __name__ == "__main__":
    main()
