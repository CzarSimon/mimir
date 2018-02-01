FROM golang:1.8-stretch as go-build

# Create source directory and bundle app source
WORKDIR /usr/src/news-ranker
COPY . /usr/src/news-ranker

# Install go dependecies
RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/httputil
RUN go get github.com/CzarSimon/go-endpoint
RUN go get github.com/CzarSimon/util

# Build go app
RUN export GOPATH=$GOPATH:$PWD
RUN go build

# python runtime
FROM python:3.6-slim-stretch as python-run

# Scikit-learn dependencies
RUN apt-get update
RUN apt-get install -y apt-utils
RUN apt-get update && apt-get install -y \
    libatlas-dev libatlas3-base libblas-dev liblapack-dev libatlas-base-dev gfortran
RUN update-alternatives --set libblas.so.3 \
    /usr/lib/atlas-base/atlas/libblas.so.3
RUN update-alternatives --set liblapack.so.3 \
    /usr/lib/atlas-base/atlas/liblapack.so.3

# Newspaper dependencies
RUN apt-get install -y libxml2-dev libxslt-dev
RUN apt-get install -y libjpeg-dev zlib1g-dev libpng-dev

# Install python dependecies
RUN pip3 install --upgrade pip
RUN pip3 install newspaper3k
RUN pip3 install requests
RUN pip3 install numpy
RUN pip3 install scipy
RUN pip3 install nltk
RUN pip3 install scikit-learn

# Corpora required to run newspapers nlp functions
RUN apt-get install -y curl
RUN curl https://raw.githubusercontent.com/codelucas/newspaper/master/download_corpora.py | python3

# Create app directory
WORKDIR /root/
COPY --from=go-build /usr/src/news-ranker/news-ranker .
COPY --from=go-build /usr/src/news-ranker/article_ranker/ article_ranker/

# Build python files
RUN python3 -m compileall -b ./article_ranker
RUN ls ./article_ranker
RUN rm ./article_ranker/*.py
RUN ls ./article_ranker

EXPOSE 5000
CMD [ "./news-ranker" ]
