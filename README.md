# tpm

A CLI client to search and retrieve passwords from [TeamPasswordManager](http://teampasswordmanager.com/).

## Getting started

Install running:

    go get github.com/martensson/tpm

Login using HMAC-key:

    tpm login

Search passwords:

    tpm search foobar

Show password by ID:

    tpm show 100

Create password:

    tpm create amazon.com --username joe --password abc123 --project 10 --email joe@test.com --tags amazon,aws,shopping
