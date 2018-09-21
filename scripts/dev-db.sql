CREATE DATABASE streamlistener;
CREATE ROLE streamlistener WITH LOGIN PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE streamlistener TO streamlistener;

CREATE DATABASE spamfilter;
CREATE ROLE spamfilter WITH LOGIN PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE spamfilter TO spamfilter;
