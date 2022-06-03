FROM postgres:14.2
LABEL maintainer="Gregory @migregal Mironov"

RUN apt-get update \
    && apt-get -y install \
        python3 \
        python3-pip \
        postgresql-plpython3-14 \
    && pip3 install tarantool\>0.4
