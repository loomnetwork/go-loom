# Backup and Restore Introduction.

A little bit of care needs to be taken to ensure that backups are usable. So it is worth testing your backups from time to time.

There are several different ways to do it, this document details one way that is known to work.

Scripts to make this easier for you can be found [here](https://github.com/loomnetwork/go-loom/tree/master/backup/utils).

# Backup and Restore Overview.

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
  * Restore the data.
  * Restore the config. Note, you may need to alter the service definition. See NODE_KEYS below.
  * Start the service.

# Backup.

* Create an extra node, which will be the backup node. It should have it's `--peers` pointing to the live nodes. They should not reference this node at all.
* Run [backup.sh](https://github.com/loomnetwork/go-loom/tree/master/backup/utils/backup.sh) periodically.

## backup.sh.

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
* STUFF_TO_BACKUP
  * For the backup node use `$BACKUP_DATA`.
  * For the other nodes, use `$BACKUP_CONFIG`.

## Considerations.


### SHOULD_SHUTDOWN.

#### Config only.

There's no need to shutdown to backup config from the live nodes.

#### Data only.

If you dig around in the settings of backup.sh, you'll see there are several different ways of configuring the service.

Do note that we've found backups taken on a machine where the serice remains running, to be unusable once the data grows beyond a certain size. It's highly recommended to stop the service to take the backup.


### BACKUP_METHOD.

If the service is running while the backup is taken, it is very likely that the data will be in an inconsistent state when the backup is taken, and will therefore be significantly more difficult to restore.

# Restore.

* Stop the service on all nodes you want to restore the data on.
* Restore the data.
* Restore the config. Note, you may need to alter the service definition. See NODE_KEYS below.
* Start the service.

## Doing the restore.

*Stop the service.*

```
  sudo systemctl stop loom.service
```

*Restore the data.*

Un-tar the data backup, restoring `app.db` and `chaindata` folders to the working directing where the loom service is run from.

*You need to manage these config files*

* genesis.json
* loom.yml - Make sure the Redis IP is correct after the restore.
* /etc/systemd/system/$SERVICE_NAME - If you restore onto different hosts/IPs, you will need to update this. See NODE_KEYS below.
* chaindata/config/node_key.json
* chaindata/config/genesis.json
* chaindata/config/config.toml

*Start everything.*

## Considerations.

### Restore order.

Depending on your configuration, you may or may not have configuration in with the data backup. To prevent accidents, it's worth restoring the data backup first, then restore the config backup after.

### NODE_KEYS.

If you are restoring backups to the same nodes that they came from, you almost certainly don't need to worry about this, because the required configuration is likely to be exactly the same. For every other situation, read on.

IMPORTANT `chaindata/config/node_key.json`, and `--peers` in `/etc/systemd/system/$SERVICE_NAME` represent the same key in different formats. Therefore when in any doubt, it's important do check as follows:

* Stop the loom service if it's running.
* Restore `chaindata/config/node_key.json` from the correct backup.
* Run `loom nodekey` from the working directory to see the configured NODE_KEY.

If you are restoring to a cluster that has previously been deployed to (Eg we set up a blank cluster using ansible, which includes setting up the NODE_KEYS). you can use the [mangleDCBackupServiceDefinition.sh](https://github.com/loomnetwork/go-loom/tree/master/backup/utils/mangleDCBackupServiceDefinition.sh) to grab the existing IPs from `/etc/systemd/system/$SERVICE_NAME`, along with the NODE_KEYS from the relevant backup to create a new `/etc/systemd/system/$SERVICE_NAME`. You will need to do this on each node, because each node excludes its own details.

## Solutions to common problems.

### Unexpected IDs, Keys, Tockens in the logs.

Almost certainly this is a miss-match between the `--peers` in `/etc/systemd/system/$SERVICE_NAME` and the values in `chaindata/config/node_key.json`.
