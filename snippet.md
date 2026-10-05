goose \
  -dir ./migrations \
  postgres \
  "postgres://orkio:password@localhost:5439/orkio?sslmode=disable" \
  up