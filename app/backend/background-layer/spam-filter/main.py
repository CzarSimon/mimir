import falcon
from handlers import ClassifyResource
from health_check import HealthCheck


api = application = falcon.API()


classifier = ClassifyResource()
health_check = HealthCheck()
api.add_route('/classify', classifier)
api.add_route('/health', health_check)
