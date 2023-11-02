# Dagger Fly.io Module

![dagger-min-version](https://img.shields.io/badge/dagger%20version-v0.9.2-green)

Manage apps on <https://fly.io>

## Deploy

Assumes that there is a valid `fly.toml` at the `--dir` path:

```sh
dagger call deploy --dir . --token "$YOUR_FLYIO_AUTH_TOKEN"
```
