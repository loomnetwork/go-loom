# Backup and Restore Introduction

A little bit of care needs to be taken to ensure that backups are usable. So it is worth testing your backups from time to time.

There are several different ways to do it, this document details one way that is known to work.

# Backup and Restore Overview

* Backup
  * On the existing nodes
    * Backup the config. This can be done with using [backup.sh](https://github.com/loomnetwork/go-loom/tree/master/backup/utils/backup.sh) and configuring `STUFF_TO_BACKUP` accordingly.
  * Create an extra node that has `--peers` for the other nodes, but is not in the `--peers` of the other nodes.
    * Run the [backup.sh](https://github.com/loomnetwork/go-loom/tree/master/backup/utils/backup.sh) script periodically.
      * Stopps the service.
      * Takes a dump.
      * Starts the service.
      * Uploads to S3.
* Restore
  * Stop the service on all nodes you want to restore the data on.
  * Restore the config. Note, you may need to alter the service definition. See below.
  * Restore the data.
  * Start the service.

# Backup

Scripts to make this easier for you can be found [here](https://github.com/loomnetwork/go-loom/tree/master/backup/utils). Of interest are

* backup.sh - For performing the actual backup.
* mangleDCBackupServiceDefinition.sh - for helping get the correct IPs and tokens into the service definition.


## backup.sh

This should be run by cron as the same user that the service runs as, or as a user with enough permissions to 

* read everything listed in the STUFF_TO_BACKUP variable.
* sudo to stop and start the service. Info for doing that [here](https://unix.stackexchange.com/questions/18830/how-to-run-a-specific-program-as-root-without-a-password-prompt).

It can be scheduled in cron like this

```
    MAILTO=""
    55 * * * * /home/ubuntu/backup.sh
```

There are several settings you can tune. Of particular interest are

* S3BUCKET - Where to send the backup.
* S3BUCKET_REGION - The region for that bucket.
* SHOULD_STOP_SERVICE - This is important. Discussed below in Considerations.


## Considerations


### Stopping the service

If the service is running while the backup is taken, it is very likely that the data will be in an inconsistent state when the backup is taken, and will therefore be significantly more difficult to restore.

Stopping the service will cause some down time, but gives a high degree of confidence that the data will be consistent.

One compromise is to back up all nodes, but only stop the service on one of those nodes. If the backup is inconsistent, you can use the data from that one node, on the remaining nodes.

# Restore

## Techniques

TODO Write this

## Solutions to common problems.

TODO Write this

## mangleDCBackupServiceDefinition.sh

If you want to restore the data to different nodes (and/or different IPs), you will need to update the IPs and tokens. This script will help you.

TODO write about when and how you'd use this.

