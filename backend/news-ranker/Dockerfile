FROM node:latest

# Install python and pip
RUN \
  apt-get update && \
  apt-get install -y python3 python3-dev python3-pip python3-setuptools && \
  rm -rf /var/lib/apt/lists/*

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
RUN pip3 install newspaper3k
RUN pip3 install requests>=2.11.1
RUN pip3 install numpy
RUN pip3 install scipy
RUN pip3 install nltk
RUN pip3 install scikit-learn

# Corpora required to run newspapers nlp functions
RUN curl https://raw.githubusercontent.com/codelucas/newspaper/master/download_corpora.py | python3

# Create app directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# Install node dependecies
COPY package.json /usr/src/app
RUN npm install

# Bundle app source
COPY . /usr/src/app

# Build python files
RUN python3 -m compileall -b ./article_ranker
RUN ls ./article_ranker
RUN rm ./article_ranker/*.py
RUN ls ./article_ranker

EXPOSE 5000
CMD [ "npm", "start" ]
