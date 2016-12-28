from sklearn.feature_extraction.text import TfidfVectorizer as tvec
from sklearn.metrics.pairwise import cosine_similarity

import news_corpus
import stopwords


def calc_reference_score(followers, previous_score=0.0):
    total_twitter_users = float(315000)
    return float(followers) / total_twitter_users + previous_score


def calc_subject_scores(subjects, text):
    query = _create_query_list_and_map(subjects)
    scores = _calc_scores(query["list"], text)
    subject_scores = {}
    for index, ticker in query["index_map"].items():
        subject_scores[ticker] = scores[index][0]
    return subject_scores


def _calc_scores(query, text):
    test_set = query + [text]
    train_set = test_set + [news_corpus.english]
    tfidf_matrix = tvec(stop_words=stopwords.english).fit_transform(train_set)
    return cosine_similarity(tfidf_matrix[0:-2], tfidf_matrix[-2:-1])


def _create_query_list_and_map(subjects):
    query_list = []
    index_map = {}
    for index, subject in enumerate(subjects):
        query_list += [subject["name"] + " " + subject["ticker"]]
        index_map[index] = subject["ticker"]
    return {"list": query_list, "index_map": index_map}
