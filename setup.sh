#!/bin/bash

echo "Please run this script as root"
tgtdir=/usr/local/bin/dns4e

echo "Creating install directory "+$tgtdir
mkdir -p $tgtdir
echo "Copying binaries into install directory "+$tgtdir
cp bin/arm/*  $tgtdir
echo "Done, please add a CRON job that is suitable using crontab -e"+$tgtdir
