[![tested-with-dagger-version](https://img.shields.io/badge/Tested%20with%20dagger-0.11.9-success?style=for-the-badge)](https://github.com/dagger/dagger/releases/tag/v0.10.1)

# Dagger Fly.io Module

![dagger-min-version](https://img.shields.io/badge/dagger%20version-v0.11.9-green)

Manage apps on <https://fly.io>

```sh
dagger functions
Name     Description
create   Creates app - required for deploy to work: `dagger call ... create --app=gostatic-example-2024-0703`
deploy   Deploys app from current directory: `dagger call ... deploy --dir=hello-static`
```

## Create

```sh
dagger call --org=personal --token=env:FLY_API_TOKEN create --app=gostatic-example-2024-07-03
```

## Deploy

Assumes there is a valid `fly.toml` at the `--dir` path:

```sh
dagger call --org=personal --token=env:FLY_API_TOKEN deploy --dir=hello-static
```
