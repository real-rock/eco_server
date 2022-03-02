FROM python:3.9
LABEL maintainer="eco@economicus.kr"
LABEL version="1.0.0"
LABEL description="Quant"

ENV PYTHONUNBUFFERED 1

RUN apt-get update && apt-get install -y wget

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN python -m pip install --upgrade pip && \
    pip install -r requirements.txt