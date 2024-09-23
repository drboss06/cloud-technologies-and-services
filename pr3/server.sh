#!/bin/sh

while true
do
  echo "-------------------"
  echo -e "HTTP/1.1 200 OK\n\n$(ls)" | nc -l -p 8080 -q 1
  request=$(nc -l -p 8080 -q 1)

  # Проверяем, создается ли файл или каталог
  if echo "$request" | grep -q "CREATE FILE"; then
    filename=$(echo "$request" | cut -d ' ' -f 3)
    touch "$filename"
    echo "Created file: $filename" | nc -l -p 8080
  elif echo "$request" | grep -q "CREATE DIR"; then
    dirname=$(echo "$request" | cut -d ' ' -f 3)
    mkdir "$dirname"
    echo "Created directory: $dirname" | nc -l -p 8080
  fi
done
