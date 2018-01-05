password=$1
name=default
remote=simon@mimir-proxy-lb
remote_path=/etc/nginx/sites-available/


# Move conf file to remote
rsync $name $remote:$proxy_path

# Test new config
ssh $remote echo $password | sudo -S nginx -t
