from sklearn.svm import LinearSVC
from sklearn.feature_extraction.text import CountVectorizer
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.pipeline import Pipeline
import numpy
import stopwords
import data_handler


def setup_model():
    return Pipeline([
        ('counter', CountVectorizer(stop_words=stopwords.english)),
        ('classifier', LinearSVC())
    ])


def train_model():
    training_data = data_handler.get_training_data()
    model = setup_model()
    model = model.fit(training_data["data"], training_data["labels"])
    return model


def classify(model, text):
    new_text = [text]
    prediction = model.predict(new_text)[0]
    return prediction
