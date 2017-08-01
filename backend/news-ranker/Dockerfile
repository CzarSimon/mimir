FROM golang:1.8 as go_builder
RUN mkdir -p /usr/src/news-ranker
WORKDIR /usr/src/news-ranker

# Bundle app source
COPY . /usr/src/news-ranker

RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/util
RUN export GOPATH=$GOPATH:$PWD
RUN go build

# python runtime
FROM python:3.6-stretch

# Scikit-learn dependencies
RUN apt-get update && apt-get install -y libatlas-dev libatlas3gf-base libblas-dev liblapack-dev libatlas-base-dev gfortran
RUN update-alternatives --set libblas.so.3 \
    /usr/lib/atlas-base/atlas/libblas.so.3
RUN update-alternatives --set liblapack.so.3 \
    /usr/lib/atlas-base/atlas/liblapack.so.3

# Newspaper dependencies
RUN apt-get install -y libxml2-dev libxslt-dev
RUN apt-get install -y libjpeg-dev zlib1g-dev libpng12-dev

# Install python dependecies
RUN pip3 install --upgrade pip
RUN pip3 install newspaper3k
RUN pip3 install requests
RUN pip3 install numpy
RUN pip3 install scipy
RUN pip3 install nltk
RUN pip3 install scikit-learn

# Corpora required to run newspapers nlp functions
RUN curl https://raw.githubusercontent.com/codelucas/newspaper/master/download_corpora.py | python3

# Create app directory
RUN mkdir -p /usr/src/news-ranker
WORKDIR /usr/src/news-ranker
COPY --from=go_builder /usr/src/news-ranker/news-ranker .
COPY --from=go_builder /usr/src/news-ranker/article_ranker/ .

# Build python files
RUN python3 -m compileall -b ./article_ranker
RUN ls ./article_ranker
RUN rm ./article_ranker/*.py
RUN ls ./article_ranker

EXPOSE 5000
CMD [ "./news-ranker" ]
