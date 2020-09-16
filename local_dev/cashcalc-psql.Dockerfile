FROM postgres:latest

COPY mock_users.sql /docker-entrypoint-initdb.d/