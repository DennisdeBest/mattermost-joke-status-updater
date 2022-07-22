#!/bin/sh
set -e

wget -c https://github.com/DennisdeBest/mattermost-joke-status-updater/releases/download/v0.0.9/mattermost-status-joke-updater-amd64 -O /usr/local/bin/mattermost-joke-status-updater
chmod +x /usr/local/bin/mattermost-joke-status-updater