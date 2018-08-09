#!/bin/bash
# Backup a Loom database with app data to S3.
# Can be scheduled like this
#   MAILTO=""
#   55 * * * * /home/ubuntu/backup.sh

# Run ./backup.sh config to dump these settings to ~/.backup_config


## BEGIN Default config ##

# Stopping gives a more confident backup, but gives down time.
SHOULD_STOP_SERVICE="false"
SERVICE_NAME="loom.service"

# Bucket upload details.
S3BUCKET="s3://your-s3-bucket/"
S3BUCKET_REGION="us-east-2"

# What we want to back up.
STUFF_TO_BACKUP="chaindata app.db genesis.json loom.yml /etc/systemd/system/$SERVICE_NAME"

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

## END Default config ##


# Where backup config is stored.
CONFIG_LOCATION=~/.backup_config

function rsyncBasedBackup
{
  # This works on the premise that stuff is going to change while we back it up,
  # so let's just cope with that and update the backup 
  
  # Get the initial backup. This is likely to have inconsistencies.
  cd "$START_DIRECTORY"
  time rsync -ru --delete $STUFF_TO_BACKUP "$TMP_LOCATION"
  
  # Repeat to get any major changes that happened during the first go.
  time rsync -ru --delete $STUFF_TO_BACKUP "$TMP_LOCATION"
  
  # Repeat to get any minor changes that happened during the second go.
  doStop
  sleep 1
  time rsync -ru --delete $STUFF_TO_BACKUP "$TMP_LOCATION"
  doStart
  
  
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

  doStop
  if [ "$SHOULD_COMPRESS" == 'true' ]; then
    tar -cjf $FILENAME -C "$START_DIRECTORY" $STUFF_TO_BACKUP
  else
    echo "Sanity check: The oldStyleBackup only does a compress, but SHOULD_COMPRESS is not true. Check configuration."
  fi
  doStart
}






# Supporting functionality.
function sanityChecks
{
  for TOOL in rsync aws; do
    if ! which "$TOOL" > /dev/null; then
      echo "Couild not find ${TOOL}. Backup will not work."
      exit 1
    fi
  done
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

function doStop
{
  if [ "$SHOULD_STOP_SERVICE" == 'true' ]; then
    sudo systemctl stop "$SERVICE_NAME"
  fi
}

function doStart
{
  if [ "$SHOULD_STOP_SERVICE" == 'true' ]; then
    sudo systemctl start "$SERVICE_NAME"
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
