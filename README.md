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

## Building and Installing Gator

### Building
If you'd like to build the aggregator tool from the source,
1. Be sure you have go install as it is a prerequisites for development and use of this application
2. Clone this repository
3. Be sure you are in the root directroy where the repo was cloned.
4. Run `go build .`

This will creating and executable binary with the name `gator`.

### Installing
If you would like to install the latest version of 'gator' simply run:
```go install github.com/mgmaster24/gator```

This will install the production binary in you Go bin directory.

## Running Gator
Now that we've discussed developing, building and installing the aggregator, we should probably learn how to use it.

### Configuration
`Gator` looks for a JSON config file in your user's home directory.  The file is name `.gatorconfig.json`. 
(Ex. `~/.gatorconfig.json`)

This file is in charge of creating the database connection to our postgres database. As well as for tracking
the currently active user.  An example is shown below.

```File: .gatorconfig.json
───────┼─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
   1   │ {
   2   │   "db_url": "postgres://postgres:my-safe-pw@localhost:5432/gator?sslmode=disable",
   3   │   "current_user_name": "someone_cool"
   4   │ }
```

You can copy the example and change the necessary pieces.  Most notable the `postgres` username and password.

### Commands

#### Login
Will login the provided username. Sets the config to use this user.

```gator login user-name```

#### Register
Register a user with the aggregator.  Sets the config to use this user.

```gator register user-name```

#### Reset
Removes all users from the database.  Mostly used for testing

```gator reset```

#### Users
Lists all the users added to the aggregator.

```gator users```

#### Agg
A continuously running method that will aggregate posts for all feeds.  A threshold in seconds
is required to tell Gator how often to fetch posts for the oldest fetched feed.

```gator agg 5s```


#### Add Feed
Adds a feed to the aggregator.

```gator addfeed name-of-feed url-of-feed```

#### Follow
Marks the currently logged in user as following a feed.  Gator will get the currently logged in user from
the gator config.

```gator follow url-of-feed```

#### Unfollow	
Removes the currently logged in user from following a feed. Gator will get the currently logged in user
from the gator config.

```gator unfollow url-of-feed```


#### Following	
List the feeds the currently logged in user is following.  Gator will get the currently logged in users
from the gator config.

```gator following```

#### Browse
Lists the posts from the feeds that the currently logged in user is following.
The parameter provided is an optional number of posts to return.  If it is not
provided the default value of 2 will be used.


```gator browse 10```

## TODO
- Add sorting and filtering options to the browse command
- Add pagination to the browse command
- Add concurrency to the agg command so that it can fetch more frequently
- Add a search command that allows for fuzzy searching of posts
- Add bookmarking or liking posts
- Add a TUI that allows you to select a post in the terminal and view it in a more readable format (either in the terminal or open in a browser)
- Add an HTTP API (and authentication/authorization) that allows other users to interact with the service remotely
- Write a service manager that keeps the agg command running in the background and restarts it if it crashes
