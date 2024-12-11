# Go Gator
Gator is a blog aggregator written in Go!  

## Prerequisites

### Download and Install Go (Golang)
In order to change and run the program you will need Go installed. 
Please view the [Go Documentation](https://go.dev/doc/install) for this step.

### Postgres
This application uses Postgres as its main database.*  You will have to install and 
configure Postgres to work on your system.  The below steps assume a Linux based
environment

1. [ ] Install Postgres v15 or later.
**Mac OS with brew**

```
brew install postgresql@15
```

**Linux / WSL (Debian).** 

```
sudo apt update
sudo apt install postgresql postgresql-contrib
```

2. [ ]   Ensure the installation worked. The psql command-line utility is the default client for Postgres. 
        Use it to make sure you're on version 15+ of Postgres:
`psql --version`

3. [ ] (Linux only) Update postgres password:
`sudo passwd postgres`
Enter a password, and be sure you won't forget it.

4. [ ] Start the Postgres server in the background
    - Mac: `brew services start postgresql`
    - Linux: `sudo service postgresql start`

5. [ ] Connect to the server. I recommend simply using the `psql` client. It's the "default" client for Postgres, 
    and it's a great way to interact with the database. While it's not as user-friendly as a GUI like PGAdmin, 
    it's a great tool to be able to do at least basic operations with.

Enter the `psql` shell:

Mac: `psql postgres`
Linux: `sudo -u postgres psql`

You should see a new prompt that looks like this:

```
postgres=#
```
## Creating the Database
If you followed the above steps you should still be in the `plsql` shell.  Now we must create our database to 
store our aggregator data.

1. [ ] Create a new database:
```CREATE DATABASE gator;```

2. [ ] Connect to the new database:
```\c gator```

You should see a new prompt that looks like this:
```gator=#```

3. [ ] Set the user password (Linux only)
```ALTER USER postgres PASSWORD 'postgres';```

I used postgres as the password.

Query the database
From here you can run SQL queries against the gator database. For example, to see the version of Postgres you're running, you can run:

```SELECT version();```

If everything is working, we can move to the next steps.
You can type `exit` to leave the psql shell.

## Development Tools
The aggregator uses some convenient tools (written in GO!) to make development a little easier.  Instructions for installing
an using these tools can be found below.

### Goose
[Goose](https://github.com/pressly/goose) is a database migration tool written in Go. 
It conveniently runs database migrations from a set of SQL files in our project

#### Install
```go install github.com/pressly/goose/v3/cmd/goose@latest```

Run `goose -version` to make sure it installed properly.

To run the necessary migrations for the aggregator:
1. [ ] Change to the `sql/schema` directory
2. [ ] Run `goose postgres postgres://<your-postgres-username>:<password>@localhost:5432/gator up`
    - If something goes wrong you can run the down migration:
        ```goose postgres postgres://<your-postres-username>:<passwor>@localhost:5432/gator down```

### SQLC
[SQLC](https://sqlc.dev/) is another convenient tool that will create Go code from our SQL queries.

**NOTE:**
You shouldn't have to make use of SQLC unless you want to create new queries.  The queries needed
for the application to function have been written.

#### Install
```go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest```

#### Configuration
SQLC is configured by the [sqlc.yaml](./sqlc.yaml) at the root of the project.  
