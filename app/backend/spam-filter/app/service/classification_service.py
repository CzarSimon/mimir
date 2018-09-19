# Internal modules
from app.config import CASHTAG_THRESHOLD
from app.models import Label, ModelType


class ClassifcationService(object):

    def __init__(self, model_repo):
        self.__repo = model_repo

    def classify(self, text, model_type=ModelType.SVM):
        """Classifes a spam candidate.

        :param text: Text to check for indications of spam.
        :param model_type: ModelType to use for classification.
        :return: Label for the tested text.
        """
        if self.__to_many_cashtags(text):
            return Label.SPAM.value
        model = self.__repo.get_spam_classifier(model_type)
        return model.predict([text])[0]

    def has_model(self):
        """Checks if the services has a trained model.

        :return: Boolen indication if the service has a trained model.
        """
        svm_model = self.__repo.get_spam_classifier(ModelType.SVM)
        nb_model = self.__repo.get_spam_classifier(ModelType.NAIVE_BAYES)
        return svm_model != None and nb_model != None

    def __to_many_cashtags(self, text):
        """Check if a given text has to high a cashtag ratio,
        messured as the percentatge of words that are cashtags.

        :param text: Text to test.
        :return: Boolean.
        """
        words = text.split()
        cashtags = [word for word in words if word.startswith('$')]
        return (len(cashtags) / len(words)) > CASHTAG_THRESHOLD
