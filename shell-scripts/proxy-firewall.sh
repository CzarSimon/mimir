proxy_ip="1.1.1.1"
ports=("3000", "5050", "7000")
for port in "${ports[@]}"
do
  ufw allow from $proxy_ip $port/tcp
done
