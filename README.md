# gostarted
Application for registration, receipt, updates tasks timedatas

## Action before start
1. Copy "config.json.default" as "config.json" and edit values
2. Run "docker-compose build"
3. Attach to container "api" and run "sql-migrate up -env='production'" to start the migration process
4. Start application "docker-compose up api"
5. Start application "docker-compose up postgres-goapp"


## Edit pg_hba.conf
cat /var/lib/postgresql/data/pg_hba.conf
ALTER ROLE postgres WITH PASSWORD 'qwerty';
sql-migrate up -env='production'

## Use sql-migrate
dbconfig.yml - settings for sql-migrate
### Create migration
/home/jarmi/go/bin/sql-migrate new "Create timeframes table"
### Up all migrations with production DB settings
sql-migrate up -env='production'
### Down 1 migration with production DB settings
sql-migrate down -env='production'
