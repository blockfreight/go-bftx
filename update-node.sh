#!/bin/bash

index=0
for validator in $BFTX0_MASTER_IP $BFTX1_MASTER_IP $BFTX2_MASTER_IP $BFTX3_MASTER_IP
do
    validator_name=bftx${index}

    ssh -oStrictHostKeyChecking=no $BLOCKFREIGHT_SSH_USER@$validator "env validator_name=$validator_name;echo $validator_name;curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/app.yaml > app.yaml;CONTAINER=`curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/app.yaml | grep blockfreight/go-bftx: | awk '{print $2}' | sed -e 's#.*:\(\)#\1#'`;echo $CONTAINER;sed -i -- 's|<VALIDATOR_NAME>|'$validator_name'|g' app.yaml;sed -i -- 's/'$CONTAINER'/'$1'/g' app.yaml;kubectl apply -f app.yaml && kubectl delete pods --all --grace-period=0 --force" <<-'ENDSSH'
        
ENDSSH
    ((index+=1))
done

# ssh -oStrictHostKeyChecking=no $BLOCKFREIGHT_SSH_USER@$validator "env validator_name=$validator_name;echo $validator_name;curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/resouces/blockfreight-statefulset.yaml > app.yaml;CONTAINER=`cat app.yaml | grep blockfreight/go-bftx: | awk '{print $2}' | sed -e 's#.*:\(\)#\1#'`;sed -i -- 's/<VALIDATOR_NAME>/'$validator_name'/g' *;sed -i -- 's/$CONTAINER/'$1'/g' *;kubectl apply -f app.yaml && kubectl delete pods --all --grace-period=0 --force" <<-'ENDSSH'
# ssh -oStrictHostKeyChecking=no $BLOCKFREIGHT_SSH_USER@$validator "env validator_name=$validator_name;echo $validator_name;curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/resouces/blockfreight-statefulset.yaml > statefulset.yaml;python3 /home/$1/image_modifier.py $2 $3;kubectl apply -f statefulset.yaml && kubectl delete pods --all --grace-period=0 --force" <<-'ENDSSH'

sed: can\'t read sedf2sN4V: Permission denied
sed: -e expression #1, char 0: no previous regular expression
