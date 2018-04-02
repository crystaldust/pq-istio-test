#!/bin/bash
hostip=$1
imagename=$2
imagefilename=$(echo $imagename | sed 's/\//_/g')

echo extracting $imagename
ssh root@$hostip "/usr/local/bin/docker save ${imagename} > /tmp/${imagefilename}"
scp root@${hostip}:/tmp/${imagefilename} /tmp/${imagefilename}

docker load < /tmp/${imagefilename}
ssh root@$hostip "rm /tmp/${imagefilename}"
