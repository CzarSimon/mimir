FROM python:2

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
