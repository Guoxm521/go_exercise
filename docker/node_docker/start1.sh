#/bin/bash
containerList=`docker ps -a|awk '{print $1}'`
echo ${containerList[@]}
# for containerID in ${containerList[@]}
# do
#          docker stop $containerID & docker rm $containerID
# done

demo=`echo hello the world | awk '{print  $1}'`
echo ${demo[@]}
for v in ${demo[@]}
do
         echo ${v}
done
