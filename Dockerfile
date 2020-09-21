FROM postgres

EXPOSE 5432

COPY ./init.sql /docker-entrypoint-initdb.d/
