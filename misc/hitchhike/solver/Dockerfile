FROM python:3.9-alpine

WORKDIR /app
ADD . /app

RUN apk add gcc musl-dev
RUN pip install -r requirements.txt

CMD ["python3", "/app/solver.py"]
