FROM python:3.8-buster

WORKDIR /app
ADD . /app

RUN pip install -r /app/requirements.txt

CMD ["python3", "/app/solver.py"]
