# Mattermost custom status Joke updater

## Install latest version

```shell
curl https://raw.githubusercontent.com/dennisdebest/mattermost-joke-status-updater/master/install.sh | sudo bash 
```

## Trigger status update 

The Mattermost API Url and secret are needed, you can pass these as parameters :

```shell
mattermost-joke-status-updater --secret=[MatterMost Secret] --url=["MatterMost Url"]
```

Or set the `MATTERMOST_URL` and `MATTERMOST_SECRET` env variables.

### Parameters

There are a few more non required parameters:
- verbose : Get more information when running the program
- name : set the specific name of the api you want to be called, if not this will be random, (options: yomomma chuck-noris jokeapi-single bread)
- maxRetries : The API will be called this many times before failing if no joke of less than 100 characters has been found (default: 10)

```shell
mattermost-joke-status-updater --secret=[MatterMost Secret] --url=["MatterMost Url"]  --name=bread --maxTries=5 --verbose 
Api name : bread 
"I'm on a roll!"
```