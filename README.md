# tpm

A CLI client to search and manage passwords from [TeamPasswordManager](http://teampasswordmanager.com/).

**Features:**

* Search and show passwords inside TeamPasswordManager.
* Support to add new passwords easily.
* Supports HMAC for improved API security.
* Single binary with no other dependencies.

## Compability

* TeamPasswordManager API v4

## Getting started

Install from source running:

    go get github.com/martensson/tpm

### Using tpm

Login using HMAC-key:

    tpm login

Search passwords:

    tpm search foobar

Show password by ID:

    tpm show 100

Create password:

    tpm create amazon.com --username joe --password abc123 --project 10 --email joe@test.com --tags amazon,aws,shopping
