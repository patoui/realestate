# Real Estate

Used as a playground to learn Go lang. Setup is heavily inspired by an [article](https://blog.logrocket.com/how-to-build-a-restful-api-with-docker-postgresql-and-go-chi/) by Michael Okoko.

#### Requirements

- [Migrate](https://github.com/golang-migrate/migrate) must be installed on the host machine
- [Make](https://www.tutorialspoint.com/unix_commands/make.htm) is installed (usually via `apt install build-essential`)


## Installation

Start by cloning the repository:

```bash
git clone git@github.com:patoui/realestate.git

OR 

git clone https://github.com/patoui/realestate.git

OR

gh repo clone patoui/realestate
```

Then make a copy of the `.env.example` file:

```bash
cp .env.example .env
```

Now start the application:

```bash
make start
```

If successfully started, visit `localhost:8080` to see the home page

To stop the docker containers, run:

```bash
make stop
```

### Bundling Frontend Assets

To bundle the frontend assets, run:

```bash
make bundle
```

### Database Migrations

To create a migration, run:

```bash
make migrate-create name=migration_name_here

# example
make migrate-create name=create_listings_table
```

To run the up migrations, run:

```bash
make migrate-up
```

And for the down migrations, run:

```bash
make migrate-down
```

### Accessing Containers

To access the go/server container, run:

```bash
make cli-server
```

To access the database/postgres container, run:

```bash
make cli-database
```

### List Commands/Help

To see CLI available commands, run:

```bash
make help
```
