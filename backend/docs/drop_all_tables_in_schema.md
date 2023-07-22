```sql
DO $$
DECLARE
  row record;
BEGIN
    FOR row IN SELECT * FROM pg_tables WHERE schemaname = 'api' 
    LOOP
      EXECUTE 'DROP TABLE api.' || quote_ident(row.tablename) || ' CASCADE';
    END LOOP;
END;
$$;
```