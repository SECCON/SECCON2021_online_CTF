FROM sagemath/sagemath:9.0

WORKDIR /app
ADD . /app

RUN sudo -S sage -pip install pycryptodome

CMD ["sudo", "-S", "sage", "problem.sage"]
