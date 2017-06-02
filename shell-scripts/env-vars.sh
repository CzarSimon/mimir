add_value() {
  echo "export $1=$2" >> ~/.bash_profile
}

enter_variable() {
  echo $1
  if [ grep -Fxq "$1" ~/.bash_profile ]
  then
    echo "not adding"
  else
    add_value $1 $2
  fi
}

enter_variable "PG_PASSWORD" "b3e7c15f3cb6f4f9ac5ee1d2a13c5104648a5f17"
enter_variable "PG_DBNAME" "mimirprod"
enter_variable "PG_USER" "simon"
enter_variable "mimir_u1" "139.59.159.73"
enter_variable "mimir_u2" "139.59.214.5"
enter_variable "mimir_u3" "46.101.133.154"

source ~/.bash_profile

update_image() {
  local image = $1
  docker pull $image
}
