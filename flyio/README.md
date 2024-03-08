[![tested-with-dagger-version](https://img.shields.io/badge/Tested%20with%20dagger-0.10.1-success?style=for-the-badge)](https://github.com/dagger/dagger/releases/tag/v0.10.1)

# Dagger Fly.io Module

![dagger-min-version](https://img.shields.io/badge/dagger%20version-v0.9.2-green)

Manage apps on <https://fly.io>

## Deploy

Assumes there is a valid `fly.toml` at the `--dir` path:

```sh
dagger call deploy --dir . --token "$YOUR_FLYIO_AUTH_TOKEN"
```
