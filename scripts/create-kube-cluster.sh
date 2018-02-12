kops create cluster \ 
  --name=kubernetes.mimir.news \
  --state=s3://mimir-kops-180101 \
  --zones=eu-east-1a \
  --node-count=2 --node-size=t2.micro --master-size=t2.micro \
  --dns-zone=mimir.news
