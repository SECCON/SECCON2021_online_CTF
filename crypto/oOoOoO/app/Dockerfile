FROM python:3.8-buster

WORKDIR /app

RUN apt update
RUN apt install socat -y

COPY ./requirements.txt /app/requirements.txt
RUN pip install -r /app/requirements.txt

COPY ./problem.py /app/problem.py
COPY ./flag.py /app/flag.py
# COPY ./server.bash /app/server.bash

CMD socat TCP-L:8000,fork,reuseaddr EXEC:"python problem.py"
