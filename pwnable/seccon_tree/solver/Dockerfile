FROM pwntools/pwntools

RUN sudo apt update && sudo apt upgrade -y && sudo apt install hashcash

COPY solve.py /solve.py

ENTRYPOINT ["python3", "/solve.py"]
