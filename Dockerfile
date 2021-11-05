FROM mysql:latest

RUN echo 'Copy table in docker entrypoint'

COPY example.sql /docker-entrypoint-initdb.d

# WORKDIR /dockerEnv

# CMD [ "go", "connectdb.go" ]