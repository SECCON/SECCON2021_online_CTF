FROM python:3.8-buster

WORKDIR /app

RUN apt update
RUN apt install socat -y

COPY ./requirements.txt /app/requirements.txt
RUN pip install -r /app/requirements.txt

COPY ./app.py /app/app.py
COPY ./flag.py /app/flag.py

CMD socat TCP-L:8080,fork,reuseaddr EXEC:"python3 app.py"
