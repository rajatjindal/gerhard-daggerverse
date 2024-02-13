[![tested-with-dagger-version](https://img.shields.io/badge/Tested%20with%20dagger-0.9.10-success?style=for-the-badge)](https://github.com/dagger/dagger/releases/tag/v0.9.10)

## Discord

### CLI example

Send Discord notifications via a [Webhook URL](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks):

```console
export DISCORD_WEBHOOK=<your-discord-webhook-here>

dagger call -m github.com/gerhard/daggerverse/notify \
    discord --webhook-url env:DISCORD_WEBHOOK --message 'Hi from Dagger Notify Module 👋'
```

## Future improvements

- [ ] Slack notifications
