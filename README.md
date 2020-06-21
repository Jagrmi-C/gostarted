# gostarted

## Action before start
1. Copy "config.json.default" as "config.json" and edit values
2. Run "docker-compose build"
3. Attach to container "api" and run "sql-migrate up -env='production'" to start the migration process
4. Start application "docker-compose up api"
5. Start application "docker-compose up postgres-goapp"


cat /var/lib/postgresql/data/pg_hba.conf
ALTER ROLE postgres WITH PASSWORD 'qwerty';
sql-migrate up -env='production'

example sql-migrate
/home/jarmi/go/bin/sql-migrate new "Create timeframes table"
