#!/bin/bash

index=0
for validator in $BFTX0_MASTER_IP $BFTX1_MASTER_IP $BFTX2_MASTER_IP $BFTX3_MASTER_IP
do
    validator_name=bftx${index}

    ssh -oStrictHostKeyChecking=no $BLOCKFREIGHT_SSH_USER@$validator "env validator_name=$validator_name;curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/app.yaml > app.yaml;CONTAINER=`cat app.yaml | grep blockfreight/go-bftx: | awk '{print $2}' | sed -e 's#.*:\(\)#\1#'`;sed -i -- 's/<VALIDATOR_NAME>/$validator_name/g' app.yaml && sed -i -- 's|'$CONTAINER'|'blockfreight/go-bftx:$1'|g' app.yaml;echo old_tag: $CONTAINER;echo tag: $1;cat app.yaml;kubectl apply -f app.yaml && kubectl delete pods --all --grace-period=0 --force;rm app.yaml pre_app.yaml" <<-'ENDSSH'

    
ENDSSH
    ((index+=1))
done

# ssh -oStrictHostKeyChecking=no $BLOCKFREIGHT_SSH_USER@$validator "env validator_name=$validator_name;echo $validator_name;curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/app.yaml > app.yaml;CONTAINER=`cat home/$2/app.yaml | grep blockfreight/go-bftx: | awk '{print $2}'`;sed -i -- 's|<VALIDATOR_NAME>|'$validator_name'|g' app.yaml;sed -i -- 's|'$CONTAINER'|'blockfreight/go-bftx:$1'|g' app.yaml;kubectl apply -f app.yaml && kubectl delete pods --all --grace-period=0 --force;cat home/$2/app.yam;" <<-'ENDSSH'

# ssh -oStrictHostKeyChecking=no $BLOCKFREIGHT_SSH_USER@$validator "env validator_name=$validator_name;echo $private_type;echo $private_key;curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/app.yaml > app.yaml;sed -i -- 's/<VALIDATOR_NAME>/'$validator_name'/g' *;sed -i -- 's/<PRIVATE_KEY>/'$private_key'/g' *;sed -i -- 's/<PRIVATE_TYPE>/'$private_type'/g' *;cat app.yaml; rm app.yaml;" <<-'ENDSSH'

# ssh -oStrictHostKeyChecking=no $BLOCKFREIGHT_SSH_USER@$validator "env validator_name=$validator_name;echo $validator_name;curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/resouces/blockfreight-statefulset.yaml > app.yaml;CONTAINER=`cat app.yaml | grep blockfreight/go-bftx: | awk '{print $2}' | sed -e 's#.*:\(\)#\1#'`;sed -i -- 's/<VALIDATOR_NAME>/'$validator_name'/g' *;sed -i -- 's/$CONTAINER/'$1'/g' *;kubectl apply -f app.yaml && kubectl delete pods --all --grace-period=0 --force" <<-'ENDSSH'
# ssh -oStrictHostKeyChecking=no $BLOCKFREIGHT_SSH_USER@$validator "env validator_name=$validator_name;echo $validator_name;curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/resouces/blockfreight-statefulset.yaml > statefulset.yaml;python3 /home/$1/image_modifier.py $2 $3;kubectl apply -f statefulset.yaml && kubectl delete pods --all --grace-period=0 --force" <<-'ENDSSH'


# ssh -oStrictHostKeyChecking=no $BLOCKFREIGHT_SSH_USER@$validator "env validator_name=$validator_name;echo $validator_name;curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/app.yaml > app.yaml;CONTAINER=`curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/app.yaml | grep blockfreight/go-bftx: | awk '{print $2}'`;echo $CONTAINER;sed -i -- 's|<VALIDATOR_NAME>|'$validator_name'|g' app.yaml;sed -i -- 's/'$CONTAINER'/'$1'/g' app.yaml;kubectl apply -f app.yaml && kubectl delete pods --all --grace-period=0 --force" <<-'ENDSSH'
# sed -i -- 's|'$CONTAINER'|'blockfreight/go-bftx:seanboi'|g' app.yaml

# sed: can\'t read sedf2sN4V: Permission denied
# sed: -e expression #1, char 0: no previous regular expression
