import req
import date
import time

schema = {
    "urls": [
        "string"
    ],
    "subjects": [
        {
            "ticker": "string",
            "name": "string"
        }
    ],
    "author": {
        "id": "int",
        "followerCount": "int"
    },
    "language": "english"
}


def new_rank_object(url, id, followers):
    return dict(
        urls=[url],
        subjects=[
            {"ticker": "AMZN", "name": "Amazon.com Inc."},
            {"ticker": "NFLX", "name": "Netflix Inc."}
        ],
        author=dict(
            id=id,
            followerCount=followers
        ),
        language="english"
    )


def test_rank_objects():
    url_1 = "https://stratechery.com/2017/amazons-earnings-amazon-logistics-services-netflixs-earnings/"
    url_2 = "https://stratechery.com/2017/amazons-earnings-amazon-logistics-services-netflixs-earnings/?test=1"
    url_3 = "https://stratechery.com/2017/amazons-earnings-amazon-logistics-services-netflixs-earnings/?test=2"
    return [
        new_rank_object(url_1, 1, 3000),
        new_rank_object(url_3, 3, 9000),
        new_rank_object(url_2, 2, 6000)
    ]


def main():
    url = "http://localhost:5000/api/rank-article"
    rank_objects = test_rank_objects()
    for rank_object in rank_objects:
        time.sleep(0.01)
        res = req.post_data(rank_object, url)
        print res


if __name__ == '__main__':
    main()
