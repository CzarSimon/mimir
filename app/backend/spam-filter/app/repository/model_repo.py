# Internal modules
from app import db
from app.models import ModelType


class ModelRepo(object):

    def __init__(self, svm_model, nb_model):
        self.__models = {
            ModelType.SVM: svm_model,
            ModelType.NAIVE_BAYES: nb_model
        }

    def get_spam_classifier(self, type=ModelType.SVM):
        """Gets spam classification model.

        :param type: Model type to use.
        :return: Spam classification model.
        """
        return self.__models[type]

    def save_classfier(self, classifer):
        """Saves a classfier in the database.

        :param classifier: Classifier to save
        """
        db.session.add(classifer)
        db.session.commit()
