## Shell scripting
In this homework, you will write a notifying program that monitors when someone logs in to your Linux server (or PC).
Notifications should be sent to two different destinations like _Discord_ & _Telegram_.
The logs should include the following information about each event:
* Date
* Username
* Hostname
* The source IP If it's a remote connection.

## Remarks
1. Your script should properly respond to `--help` and `-h` options.
2. You can choose between implementing a script that listens to `PAM` events or `TTY` events.
3. Make sure the script runs every time someone enters the system.
