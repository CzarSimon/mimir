import config
import data_handler
import falcon
import json
import trainer


class ClassifyResource(object):
    def __init__(self):
        self.model = trainer.train_model()
        print "resource instanziated, model trained"


    def classify(self, tweet):
        words = tweet["text"].split(" ")
        over_cashtag_threshold, text = data_handler.clean_text(words)
        if not over_cashtag_threshold:
            category = trainer.classify(self.model, text)
        else:
            category = config.categories["SPAM"]
        return _create_response(category)


    def on_get(self, request, response):
        response.status = falcon.HTTP_200
        response.body = "Hello World"


    def on_post(self, request, response):
        try:
            raw_json = request.stream.read()
        except Exception as e:
            raise falcon.HTTPError(falcon.HTTP_400, 'error', e.message)
        try:
            tweet = json.loads(raw_json, encoding='utf-8')
        except ValueError:
            raise falcon.HTTPError(falcon.HTTP_400,
                'Malformed JSON',
                'Could not decode the request body. The JSON was incorrect.')
        response.status = falcon.HTTP_200
        response.body = self.classify(tweet)


def _create_response(category):
    response = {
        "result": category
    }
    return json.dumps(response, encoding='utf-8')
