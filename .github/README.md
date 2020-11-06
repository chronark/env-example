# env-example
Create a .env.example file by removing values from a .env file.


### Install

```sh
go install github.com/chronark/env-example
```
Or choose a binary from the [releases](https://github.com/chronark/env-example/releases).

### Usage
```
example-env -in=path/to/.env -out=.env.example
```
Omiting in and out defaults to `.env` and `.env.example`

### Example
`.env`
```sh
# Comment
HELLO=World

# Empty value
NO_VALUE=
```

Will become:

```sh
# Comment
HELLO=

# Empty value
NO_VALUE=
```
