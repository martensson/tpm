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
