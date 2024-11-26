# exam files

## docker compose setup

### database

Database is served via `./backups/data_backup.sql` which applied by container `init_db` at the start of the app.<br>
It's okay when you see `init_db` failed — in most recent case that means that DB is not empty — migrations are already 
applied. If it's not ok for you, you need to:

> docker volume rm golang-gorm-postgres_postgres

The command clears docker volume, then you can apply db migrations.<br>
And this is how you make a dump:
> pg_dump -h localhost -p 6500 -U postgres -f data_backup.sql golang-gorm

### watermark

> convert -background none -fill "rgba(255,255,255,0.5)" -font Arial -pointsize 48 label:"© Your Company" watermark.png\n