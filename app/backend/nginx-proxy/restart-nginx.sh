password=$1
ssh simon@mimir-proxy-lb echo $password | sudo -S systemctl reload nginx
