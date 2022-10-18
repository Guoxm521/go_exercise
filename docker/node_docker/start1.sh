#/bin/bash
containerList=`docker ps -a|awk '{print $1}'`
echo $containerList
for containerID in ${containerList[@]}
do
         docker stop $containerID & docker rm $containerID
done
