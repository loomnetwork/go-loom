#!/bin/bash
# Mangle an original loom service definition with one from a backup, to get the correct tokens with the correct IPs.
#
# Syntax
#   $0 originalFile backupFile

originalFile="$1"
backupFile="$2"
port=46656

function showHelp
{
  grep -B 100 greppery $0 | grep '^#' | tail -n +2 | cut -b 3- | sed 's#$0#'$0'#g'
}


function getHead
{
  grep '^ExecStart' "$originalFile" | sed 's/ tcp:\/\/.*//g'
}

function getIPs
{
  grep '^ExecStart' "$originalFile" | sed 's/:/\n/g' | grep -v 46656 | cut -d@ -f2 | tail -n +2
}

function getTokens
{
  grep '^ExecStart' "$backupFile" | sed 's#tcp://#\n#g' | tail -n +2 | cut -d@ -f1
}

function getCompleteLine
{
  # 1 Head
  # 2, 3, 4 IPs
  # 5, 6, 7 Tockens
  echo "$1" tcp://$5@$2:$port,tcp://$6@$3:$port,tcp://$7@$4:$port
}

if [ "$2" == '' ] ; then
  showHelp
  exit 1
fi

line=$(getCompleteLine "`getHead`" `getIPs` `getTokens`)

sed "s#^ExecStart.*#$line#g" "$originalFile"
