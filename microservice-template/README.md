# Skeleton for a new service

This document describes how to create a new service based on this skeleton. There **MUST** be 5 directories in your service project structure:

* `/bin`: contains our compiled application binaries, ready for deployment to a production server.
* `/cmd/api`: contains the application-specific code for your API application. This will include the code for running the server, reading and writing HTTP requests, and managing authentication.
* `/internal`: contains various ancillary packages used by our API. It will contain the code for interacting with our database, doing data validation, sending emails and so on. Basically, any code which *isn’t* application-specific and can potentially be reused will live in here.
* `/remote`: contains the configuration files and setup scripts for our production server.

And there **MUST** be 5 files in your service project structure:

* `Makefile`: contains the commands for building, testing, and running the application.
* `go.mod`: contains the Go module definition for the project.
* `README.md`: contains the documentation for the project.
* `.gitignore`: contains the files and directories that Git should ignore.
* `pre-commit-config.yaml`: contains the configuration for the pre-commit hooks (follow the instructions in the pre-commit [installation](https://pre-commit.com/#installation) website to install it and then the [quickstart](https://pre-commit.com/#installation)).