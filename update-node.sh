#!/bin/bash

index=0
for validator in $BFTX0_MASTER_IP $BFTX1_MASTER_IP $BFTX2_MASTER_IP $BFTX3_MASTER_IP
do
    validator_name=bftx${index}

    ssh -oStrictHostKeyChecking=no $BLOCKFREIGHT_SSH_USER@$validator "env validator_name=$validator_name;echo $validator_name;pwd;curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/resouces/blockfreight-statefulset.yaml > statefulset.yaml;python3 /home/sean/image_modifier.py $1 $2;kubectl apply -f statefulset.yaml && kubectl delete pods --all --grace-period=0 --force" <<-'ENDSSH'
    
    
ENDSSH
    ((index+=1))
done

# ssh -oStrictHostKeyChecking=no $BLOCKFREIGHT_SSH_USER@$validator "env validator_name=$validator_name;echo $validator_name;curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/resouces/blockfreight-statefulset.yaml > statefulset.yaml;python image_modifier.py $1 $2;sed -i -- 's/<VALIDATOR_NAME>/'$validator_name'/g' *;kubectl apply -f statefulset.yaml && kubectl delete pods --all --grace-period=0 --force" <<-'ENDSSH'