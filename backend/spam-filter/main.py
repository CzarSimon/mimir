import falcon
from handlers import ClassifyResource


api = application = falcon.API()


classifier = ClassifyResource()
api.add_route('/classify', classifier)
