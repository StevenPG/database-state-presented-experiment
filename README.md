# database-state-presented-experiment
This repository is a small experiment testing the scalability of a system that requires a caller to include the current database state in it's request and how this scales horizontally.

## Structure of the experiment

There will be a server utility that will present two endpoints, a GET and a PUT. 

The PUT will push an update to the database, with a large caveat. Each PUT request
must contain the entire database worth of values.

For example, if there are 5 entries in the database and we wish to push a sixth, we will need
to provide the 5 entries to the database to verification that we are aware of the state of the system.

This will be run with increasing numbers of entries and multiple services.

The database server will receive the request and pull all entries in the database. It will then
simply verify each key provided exists in the database by mapping the presented key and value with
the database key and value.

## Test configurations

- One service pushing 10 key:value pairs into the database. For each push (synchronously), it first
retrieves the other 9 entries and builds the request.

- Two identical services pushing 20 total key:value pairs in the database. These will potentially stumble
over each other as they will be unaware of the other services attempts to update the database.

- Further configurations are TBD if there is huge deadlocking/clashing with the above configurations.

## Findings

There is an upper limit for this sort of system. 

Found Limit 1: 3 users at 50ms per request. 4 users causes heavy failures.
Found Limit 2: 9 users at 250ms per request. 10 users causes light failures.