#!/bin/bash
# Backup a Loom database with app data to S3.
# Can be scheduled like this
#   MAILTO=""
#   55 * * * * /home/ubuntu/backup.sh

# Run ./backup.sh config to dump these settings to ~/.backup_config


## BEGIN Default config ##

# Stopping gives a more confident backup, but gives down time.
SHOULD_STOP_SERVICE="false"
STOP_TIMEOUT="300"
SERVICE_NAME="loom.service"
PROCESS_NAME="loom"

# If the service doesn't come back up, give it another kick.
SHOULD_KICK_SERVICE="true"
KICK_TIMEOUT=180
MAXIMUM_NUMBER_OF_KICKS=100
PORT="46657"

# Send a SIGUSR1 to the service.
SHOULD_SIGUSR1_SERVICE="false"
SIGUSR1_WAIT_SECONDS=3

# Bucket upload details.
S3BUCKET="s3://your-s3-bucket/"
S3BUCKET_REGION="us-east-2"

# What we want to back up.
BACKUP_DATA="chaindata app.db"
BACKUP_CONFIG="genesis.json loom.yml /etc/systemd/system/$SERVICE_NAME"
STUFF_TO_BACKUP="$BACKUP_DATA $BACKUP_CONFIG"

# Anything that doesn't have an explicit path will be backed up from this directory.
START_DIRECTORY="/home/ubuntu"

# Clean up the tmp location. Setting to false will give a faster backup at the expense of space.
SHOULD_CLEAN_TMP="false"

# How the backup is done.
# * rsyncBasedBackup - Recommended. Most likely to give reliable results.
# * oldStyleBackup - Simpler, requires less space.
BACKUP_METHOD="rsyncBasedBackup"

# Output.
# Compress. Only turn this off for testing.
SHOULD_COMPRESS="true"
SHOULD_UPLOAD="true"

# How to name the backup.
FILENAME="/tmp/`date +%Y-%m-%d--%H%M%S`-`hostname`.tar.bz2"

# Where to store the temporary files.
TMP_LOCATION=/tmp/loom-bu-stage

# Shutdown the server after a successful backup. This is useful if you have created a temporary server to replicate, backup, shutdown.
SHOULD_SHUTDOWN='false'

## END Default config ##


# Where backup config is stored.
CONFIG_LOCATION=~/.backup_config

function rsyncBasedBackup
{
  # This works on the premise that stuff is going to change while we back it up,
  # so let's just cope with that and update the backup 
  
  # Get the initial backup. This is likely to have inconsistencies.
  cd "$START_DIRECTORY"
  time rsync -rui --delete $STUFF_TO_BACKUP "$TMP_LOCATION"
  
  # Repeat to get any major changes that happened during the first go.
  time rsync -rui --delete $STUFF_TO_BACKUP "$TMP_LOCATION"
  
  # Repeat to get any minor changes that happened during the second go.
  doSIGUSR1
  doStop
  sleep 1
  time rsync -rui --delete $STUFF_TO_BACKUP "$TMP_LOCATION"
  doStart
  START_TIME=`now`
  
  # Compress it.
  if [ "$SHOULD_COMPRESS" == 'true' ]; then
    cd "$TMP_LOCATION"
    tar -cjf $FILENAME *
  fi
}

function oldStyleBackup
{
  # This works well enough if the service is stopped, but it causes some down-
  # time, and a risk that the service might not come back up with repeated
  # backups.
  # If the service isn't stopped, the backup is usually corrupted on a busy
  # server. Therefore this shouldn't be relied on.

  doSIGUSR1
  doStop
  if [ "$SHOULD_COMPRESS" == 'true' ]; then
    tar -cjf $FILENAME -C "$START_DIRECTORY" $STUFF_TO_BACKUP
  else
    echo "Sanity check: The oldStyleBackup only does a compress, but SHOULD_COMPRESS is not true. Check configuration."
  fi
  doStart
  START_TIME=`now`
}






# Supporting functionality.
function sanityChecks
{
  for TOOL in rsync aws; do
    if ! which "$TOOL" > /dev/null; then
      echo "ERROR: Couild not find ${TOOL}. Backup will not work." >&2
      exit 1
    fi
  done
  
  if [ "$SHOULD_STOP_SERVICE" == 'true' ] && [ "$SHOULD_SIGUSR1_SERVICE" == 'true' ]; then
    echo "WARN: Both SHOULD_STOP_SERVICE and SHOULD_SIGUSR1_SERVICE are configured. You probably don't want this." >&2
  fi
  
  if [ "$SHOULD_SHUTDOWN" == 'true' ] && [ "$SHOULD_KICK_SERVICE" == 'true' ]; then
    echo "WARN: Both SHOULD_SHUTDOWN and SHOULD_KICK_SERVICE are configured. You probably don't want this." >&2
  fi
}

function prep
{
  mkdir -p "$TMP_LOCATION"
}

function cleanup
{
  if [ "$SHOULD_CLEAN_TMP" == 'true' ]; then
    rm -Rf "$TMP_LOCATION"
  fi
  
  if [ "$SHOULD_UPLOAD" == 'true' ] && [ "$SHOULD_COMPRESS" == 'true' ]; then
    rm $FILENAME
  else
    echo "Upload and/or compression is turned off. Therefore won't delete. Unless If you aren't testing, this is probably a configuration issue."
  fi
}

function now
{
  date +%s
}

function doStop
{
  if [ "$SHOULD_STOP_SERVICE" == 'true' ]; then
    STOP_BEGIN=`now`
    sudo systemctl stop "$SERVICE_NAME"
    
    if pidof "$PROCESS_NAME" > /dev/null; then
      echo "Waiting for the process $PROCESS_NAME to stop."
      while pidof "$PROCESS_NAME" > /dev/null; do
        NOW=`now`
        let DURATION=$NOW-$STOP_BEGIN
        
        if [ $DURATION -gt $STOP_TIMEOUT ]; then
          echo "Timed out while stopping the service. Not sane to consinue. Here is the status of the service." >&2
          sudo systemctl status "$SERVICE_NAME" >&2
          exit 1
        fi
        sleep 0.1
      done
      echo "Process $PROCESS_NAME stopped."
    fi
  fi
}

function doStart
{
  if [ "$SHOULD_STOP_SERVICE" == 'true' ] && [ "SHOULD_SHUTDOWN" == 'false' ]; then
    sudo systemctl start "$SERVICE_NAME"
  fi
}

function kickServiceIfNotHappy
{
  if [ "$SHOULD_KICK_SERVICE" == 'true' ]; then
    if [ "$1" == '' ]; then
      TIMEOUT_START=`now`
    else
      TIMEOUT_START="$1"
    fi
    
    if ! serviceIsAlive; then
      echo "WARN: The service is not healthy. Waiting for $KICK_TIMEOUT seconds from $TIMEOUT_START to restart the service. Now=`now`."
      waitForSeconds $KICK_TIMEOUT $TIMEOUT_START
      sudo systemctl restart "$SERVICE_NAME"
      return 1
    else
      echo "INFO: The service is alive."
      return 0
    fi
  fi
}

function repetitivelyKickServiceUntilHappy
{
  if [ "$1" == '' ]; then
    TIMEOUT_START=`now`
  else
    TIMEOUT_START="$1"
  fi
  
  kickServiceIfNotHappy "$TIMEOUT_START"
  
  KICKS=1
  while ! serviceIsAlive && [ $KICKS -lt $MAXIMUM_NUMBER_OF_KICKS ] ; do
    kickServiceIfNotHappy
    let KICKS=$KICKS+1
    
    if ! [ $KICKS -lt $MAXIMUM_NUMBER_OF_KICKS ]; then
      echo "ERROR: Kicked the service $KICKS times and was not able to get it started." >&2
    fi
  done
}

function waitForSeconds
{
  WAIT_DURATION="$1"
  
  if [ "$1" == '' ]; then
    WAIT_START=`now`
  else
    WAIT_START="$2"
  fi
  
  let WAIT_CURRENT=`now`-$WAIT_START
  while [ $WAIT_DURATION -lt -$WAIT_CURRENT ]; do
    sleep 0.5
    let WAIT_CURRENT=`now`-$WAIT_START
  done
}

function serviceIsAlive
{
  curl -s localhost:$PORT/status | grep -q latest_block_height
  return $?
}

function doSIGUSR1
{
  if [ "$SHOULD_SIGUSR1_SERVICE" == 'true' ]; then
    sudo killall -SIGUSR1 loom
    sleep $SIGUSR1_WAIT_SECONDS # Give a few seconds for stuff to be written to file before we proceed with the next step.
  fi
}

function upload
{
  if [ "$SHOULD_UPLOAD" == 'true' ] && [ "$SHOULD_COMPRESS" == 'true' ]; then
    /usr/bin/aws s3 cp $FILENAME "$S3BUCKET" --region "$S3BUCKET_REGION"
  else
    echo "Upload and/or compression is turned off. Therefore won't upload. Unless If you aren't testing, this is probably a configuration issue."
  fi
}

function shutdownIfRequested
{
  if [ "SHOULD_SHUTDOWN" == 'true' ]; then
    sudo shutdown -h 0
  fi
}

function generateConfig
{
  grep -A 1000 '^## BEGIN Default config ##' $0 | grep -B 1000 "^## END Default config ##" | grep -v 'Default config ##' > "$CONFIG_LOCATION"
  echo "Default config dumped to $CONFIG_LOCATION. Edit this to get the behavior you want."
}

function doIt
{
  sanityChecks
  prep


  case "$BACKUP_METHOD" in
    "rsyncBasedBackup")
      rsyncBasedBackup
    ;;
    "oldStyleBackup")
      oldStyleBackup
    ;;
    *)
      echo "Unknown backup method."
    ;;
  esac

  upload
  cleanup
  
  # This should never be needed, but it's here just incase.
  repetitivelyKickServiceUntilHappy "$START_TIME"
}

if [ "$1" == 'config' ]; then
  generateConfig
  exit 0
fi

if [ -e "$CONFIG_LOCATION" ]; then
  . "$CONFIG_LOCATION"
else
  echo "WARNING: $CONFIG_LOCATION does not exist, so internal defaults have been used. Run ./backup.sh config to create it."
fi

doIt
