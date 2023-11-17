#!/bin/bash

TELEGRAM_BOT_TOKEN="YOUR_TELEGRAM_BOT_TOKEN"
TELEGRAM_CHAT_ID="YOUR_TELEGRAM_CHAT_ID"

tail -n 0 -F /var/log/audit/audit.log | \
  while read -r line; do
    if [[ $line == *'msg=audit('*'login='* ]]; then
      DATE=$(date +'%Y-%m-%d %H:%M:%S')
      USERNAME=$(echo "$line" | grep -oP 'acct="[^"]+"' | cut -d'"' -f2)
      HOSTNAME=$(echo "$line" | grep -oP 'hostname=[^ ]+' | cut -d'=' -f2)
      IP=$(echo "$line" | grep -oP 'addr=[^ ]+' | cut -d'=' -f2)
      
      # Format the notification message
      MESSAGE="Login Event\nDate: $DATE\nUsername: $USERNAME\nHostname: $HOSTNAME\nSource IP: $IP"
      
      # Send to Telegram
      curl -s "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/sendMessage" \
        -d "chat_id=$TELEGRAM_CHAT_ID" \
        -d "text=$MESSAGE"
    fi
  done

