import falcon


class HealthCheck(object):
    def __init__(self):
        pass


    def on_get(self, request, response):
        response.status = falcon.HTTP_200
        response.body = "OK\n"
