# Scripts

This folder contains scripts for automating often-used, multiple step tasks.


## Check Documentation

[check-docs.sh](./check-docs.sh) - This script checks if the commands documentation was changed and lists the changes.
 
 
##  CLI Installation

[install.sh](./install.sh) - This script aids in the installation of RHOAS CLI.


Anything in [./util](./util) can be considered non-essential helper scripts.

- [delete-user-kafkas.sh](./util/delete-user-kafkas.sh) - This script will delete all Kafkas for the currently logged in user. Not useful for non-TTY automation - requires human interaction for confirming the name of the Kafkas you want to delete.
- [delete-service-accounts.sh](./util/delete-service-accounts.sh) - Delete a filtered subset of all service accounts in your organisation. Running `./scripts/util/delete-service-accounts "your-name"` will delete all service accounts which have "your-name" in the name.
