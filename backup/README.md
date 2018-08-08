# Backup

A little bit of care needs to be taken to ensure that backups are usable. So it is worth testing your backups from time to time.

Scripts to make this easier for you can be found [here](https://github.com/loomnetwork/go-loom/tree/master/backup). Of interest are

* backup.sh - For performing the actual backup.
* mangleDCBackupServiceDefinition.sh - for helping get the correct IPs and tokens into the service definition.

## backup.sh

This should be run by cron as the same user that the service runs as, or as a user with enough permissions to read everything listed in the STUFF_TO_BACKUP variable.

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

