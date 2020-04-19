FROM mongo:latest

EXPOSE 27017

COPY cashcalc-db-dump cashcalc-db-dump
COPY cashcalc_startup_db.sh /docker-entrypoint-initdb.d

CMD ["mongod"]
