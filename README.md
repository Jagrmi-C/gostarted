# gostarted

cat /var/lib/postgresql/data/pg_hba.conf
ALTER ROLE postgres WITH PASSWORD 'qwerty';
sql-migrate up -env='production'

example sql-migrate
/home/jarmi/go/bin/sql-migrate new "Create timeframes table"
