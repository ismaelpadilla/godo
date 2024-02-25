# Task reminder

Set reminders and get notifications.

## Usage

```
# godo [REMINDER] @@ [TIME]
godo Remind me to do something @@ tomorrow
```

You can check current reminders and trigger notifications with:
```
godo --check
```

To automatically check for reminders, add the following to your crontab:
```
* * * * * godo --check
```

## To do

- Improve storage solution.
- Improve parser, ideally support natural language.
- Daemonize tasks checker.
