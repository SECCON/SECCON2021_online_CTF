FROM python:3.8-buster

WORKDIR /app
ADD . /app

RUN pip install -r requirements.txt

CMD ["python3", "/app/solver.py"]
