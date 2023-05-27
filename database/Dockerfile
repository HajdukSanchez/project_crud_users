# Use this Postgres image
FROM postgres:10.3

# Copy up.sql file into speicified directory to init container DB
COPY up.sql /docker-entrypoint-initdb.d/1.sql

# Run initial command
CMD ["postgres"]