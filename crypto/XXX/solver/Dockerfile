FROM sagemath/sagemath:latest

WORKDIR /
USER root

ADD ./solve.sage solve.sage

ENTRYPOINT sage solve.sage
