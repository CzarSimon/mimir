# Standard libarary
import logging

# 3rd party modules
from sklearn.pipeline import Pipeline

# Internal modules
from app import db
from app.models import Classifier, ModelType


class ModelRepo:

    __log = logging.getLogger('ModelRepo')

    def __init__(self, svm_model: ModelType, nb_model: ModelType) -> None:
        self.__models = {
            ModelType.SVM: svm_model,
            ModelType.NAIVE_BAYES: nb_model
        }

    def get_spam_classifier(self, type: ModelType = ModelType.SVM) -> Pipeline:
        """Gets spam classification model.

        :param type: Model type to use.
        :return: Spam classification model.
        """
        return self.__models[type]

    def save_classfier(self, classifer: Classifier) -> None:
        """Saves a classfier in the database.

        :param classifier: Classifier to save
        """
        if self.__classifier_exists(classifer):
            self.__log.info('Classifier already exists')
            return
        db.session.add(classifer)
        db.session.commit()

    def __classifier_exists(self, classifer: Classifier) -> bool:
        existing_classifiers = Classifier.query.\
            filter(Classifier.model_hash == classifer.model_hash).all()
        return len(existing_classifiers) > 0
