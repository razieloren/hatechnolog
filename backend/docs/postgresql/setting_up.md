```bash
pg_ctl -D /usr/local/var/postgres [start,stop,restart]
```

```bash
# Connect to default database:
psql postgres
# Create backend DB:
CREATE database hatechnolog;
\connect hatechnolog
# Schema for query module:
CREATE schema query
# Schema for REST API:
CREATE schema api
# Backend user:
CREATE USER hatechnolog_admin WITH ENCRYPTED PASSWORD 'mypass';
GRANT ALL PRIVILEGES ON DATABASE hatechnolog TO hatechnolog_admin;
GRANT ALL ON SCHEMA api,query TO hatechnolog_admin;
GRANT ALL ON ALL TABLES IN SCHEMA api,query TO hatechnolog_admin;
GRANT ALL ON ALL SEQUENCES IN SCHEMA api,query TO hatechnolog_admin;
GRANT ALL ON ALL FUNCTIONS IN SCHEMA api,query TO hatechnolog_admin;
# Frontend user:
CREATE USER hatechnolog_guest WITH ENCRYPTED PASSWORD 'mypass';
GRANT CONNECT ON DATABASE hatechnolog TO hatechnolog_guest;
GRANT USAGE ON SCHEMA api,query TO hatechnolog_guest;
GRANT SELECT ON ALL TABLES IN SCHEMA query TO hatechnolog_guest;
GRANT ALL ON ALL TABLES IN SCHEMA api TO hatechnolog_guest;
GRANT ALL ON ALL SEQUENCES IN SCHEMA api TO hatechnolog_guest;
GRANT ALL ON ALL FUNCTIONS IN SCHEMA api TO hatechnolog_guest;

# Check permissions:
\l
\dn+
```