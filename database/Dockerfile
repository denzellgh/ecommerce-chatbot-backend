FROM postgres:10.3

COPY 01-init.sql /docker-entrypoint-initdb.d/01-init.sql
COPY 02-seed.sql /docker-entrypoint-initdb.d/02-seed.sql
COPY 03-auth-schema.sql /docker-entrypoint-initdb.d/03-auth-schema.sql
#COPY 04-auth-seed.sql /docker-entrypoint-initdb.d/04-auth-seed.sql

CMD ["postgres"]