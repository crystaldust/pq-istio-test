#!/bin/bash
ipsubnet=$1
imagename=$2
imagefilename=$(echo $imagename | sed 's/\//_/g')

echo distributing $imagename
echo saving to $imagefilename.tar

docker save $imagename > /tmp/$imagefilename.tar

scp /tmp/$imagefilename.tar root@192.168.${ipsubnet}0:/root/
scp /tmp/$imagefilename.tar root@192.168.${ipsubnet}1:/root/
scp /tmp/$imagefilename.tar root@192.168.${ipsubnet}2:/root/

ssh root@192.168.${ipsubnet}0 "/usr/local/bin/docker load < /root/$imagefilename.tar"
ssh root@192.168.${ipsubnet}1 "/usr/local/bin/docker load < /root/$imagefilename.tar"
ssh root@192.168.${ipsubnet}2 "/usr/local/bin/docker load < /root/$imagefilename.tar"

rm /tmp/$imagefilename.tar
ssh root@192.168.${ipsubnet}0 "rm /root/$imagefilename.tar"
ssh root@192.168.${ipsubnet}1 "rm /root/$imagefilename.tar"
ssh root@192.168.${ipsubnet}2 "rm /root/$imagefilename.tar"
