#sample table creation at 'example' DB
migrate -database ${DATABASE_URL} -path ./migrations up

#sample table drop at 'example' DB
migrate -database ${DATABASE_URL} -path ./migrations down