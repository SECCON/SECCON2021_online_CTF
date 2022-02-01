FROM sagemath/sagemath:9.0

WORKDIR /app
ADD . /app

# CMD ["sleep", "infinity"]
# CMD ["sudo", "-S", "SECCON_HOST=$SECCON_HOST", "SECCON_PORT=$SECCON_PORT", "sage", "solver.sage"]
CMD ["sudo", "-S", "SECCON_HOST=$SECCON_HOST", "SECCON_PORT=$SECCON_PORT", "bash", "solver.bash"]
