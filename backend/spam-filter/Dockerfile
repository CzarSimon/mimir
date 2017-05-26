FROM python:2

# Scikit-learn dependencies
RUN apt-get update && apt-get install -y libatlas-dev libatlas3gf-base libblas-dev liblapack-dev libatlas-base-dev gfortran
RUN update-alternatives --set libblas.so.3 \
    /usr/lib/atlas-base/atlas/libblas.so.3
RUN update-alternatives --set liblapack.so.3 \
    /usr/lib/atlas-base/atlas/liblapack.so.3


# Create app directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# Bundle app source
COPY . /usr/src/app

# Install python dependecies
RUN pip install gunicorn==19.7.1
RUN pip install falcon==1.2.0
RUN pip install psycopg2==2.6.1
RUN pip install numpy==1.11.1
RUN pip install nltk==3.2.1
RUN pip install scipy==0.18.0
RUN pip install scikit-learn==0.17.1

EXPOSE 1000
CMD [ "gunicorn", "-w 4", "-b :1000", "main" ]
