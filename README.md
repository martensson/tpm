# tpm [![Build Status](https://travis-ci.org/martensson/tpm.svg?branch=master)](https://travis-ci.org/martensson/tpm)

A CLI client to search and manage passwords inside [TeamPasswordManager](http://teampasswordmanager.com/).

**Features:**

* Search and show passwords.
* Support to add new passwords easily.
* Supports HMAC for improved API security.
* Single binary with no other dependencies.

## Compability

* TeamPasswordManager API v4

## Installation

1. Install pre-compiled binary from [releases](https://github.com/martensson/tpm/releases).

2. Install from source: `go get github.com/martensson/tpm`

It's that easy!

### Using tpm

Login using HMAC-key:

    tpm login

Search passwords:

    tpm search aws.amazon.com

Show password by ID:

    tpm show 100

Create password:

    tpm create aws.amazon.com --username joe --password abc123 --project 10 --email joe@test.com --tags amazon,aws,shopping
