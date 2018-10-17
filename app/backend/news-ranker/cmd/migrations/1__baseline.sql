-- +migrate Up
CREATE TABLE article (
  id VARCHAR(50) PRIMARY KEY,
  url VARCHAR(350) NOT NULL,
  title VARCHAR(350) NOT NULL,
  body TEXT,
  keywords TEXT,
  reference_score NUMERIC(9,5) NOT NULL,
  article_date DATE,
  created_at TIMESTAMP
);

CREATE TABLE twitter_references (
  id VARCHAR(50) PRIMARY KEY,
  twitter_author VARCHAR(50),
  follower_count INT,
  article_id REFERENCES article(id)
)

CREATE TABLE compound_score (
  id VARCHAR(50) PRIMARY KEY,
  article_id REFERENCES article(id),
  symbol VARCHAR(15),
  score NUMERIC(9,5) NOT NULL
);

CREATE TABLE subject_score (
  id VARCHAR(50) PRIMARY KEY,
  article_id REFERENCES article(id),
  symbol VARCHAR(15),
  score NUMERIC(9,5) NOT NULL
);

CREATE TABLE article_cluster (
  cluster_hash VARCHAR(64) PRIMARY KEY,
  title VARCHAR(255),
  symbol VARCHAR(15),
  article_date DATE,
  lead_article_id VARCHAR(50) REFERENCES article(id),
  score NUMERIC(9,5),
);

CREATE TABLE cluster_member (
  id VARCHAR(50) PRIMARY KEY,
  cluster_hash VARCHAR(64) REFERENCES article_cluster(cluster_hash),
  article_id VARCHAR(50) REFERENCES article(id),
  reference_score NUMERIC(9,5),
  subject_score NUMERIC(9,5),
);

-- +migrate Down
DROP TABLE IF EXISTS cluster_member;
DROP TABLE IF EXISTS article_cluster;
DROP TABLE IF EXISTS subject_score;
DROP TABLE IF EXISTS compound_score;
DROP TABLE IF EXISTS twitter_references;
DROP TABLE IF EXISTS article;
