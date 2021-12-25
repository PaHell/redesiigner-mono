#!/bin/bash
echo "Running script from redesiigner/mysql/init.sh"

IFS=',' read -ra colonTuples <<< $MYSQL_DATABASE_ACCOUNTS
for tuple in "${colonTuples[@]}"; do
      echo "Creating user and owned database for: ${tupleSplit[0]}"
      IFS=':' read -ra tupleSplit <<< $tuple
      mysql -u root "-p${MYSQL_ROOT_PASSWORD}" -e "CREATE USER ${tupleSplit[0]}@localhost IDENTIFIED BY \"${tupleSplit[1]}\";"
      mysql -u root "-p${MYSQL_ROOT_PASSWORD}" -e "CREATE DATABASE ${tupleSplit[0]};"
      mysql -u root "-p${MYSQL_ROOT_PASSWORD}" -e "GRANT ALL PRIVILEGES ON ${tupleSplit[0]}.* TO ${tupleSplit[0]}@localhost;"
      mysql -u root "-p${MYSQL_ROOT_PASSWORD}" -e "FLUSH PRIVILEGES;"
done

