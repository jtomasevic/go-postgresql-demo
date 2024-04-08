FROM postgres:latest
ENV POSTGRES_PASSWORD imdb_user
ENV POSTGRES_USER imdb_user
ENV POSTGRES_DB imdb
COPY ./init/schema.sql /docker-entrypoint-initdb.d/
USER postgres