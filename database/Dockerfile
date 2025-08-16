FROM postgres:10.3

COPY 01-init.sql /docker-entrypoint-initdb.d/01-init.sql
COPY 02-seed.sql /docker-entrypoint-initdb.d/02-seed.sql

CMD ["postgres"]