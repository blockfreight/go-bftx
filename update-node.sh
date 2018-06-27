#!/bin/bash

CONTAINER=`curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/app.yaml | grep blockfreight/go-bftx: | awk '{print $2}' | sed -e 's#.*:\(\)#\1#'`

index=0
for validator in $BFTX0_MASTER_IP $BFTX1_MASTER_IP $BFTX2_MASTER_IP $BFTX3_MASTER_IP
do
    validator_name=bftx${index}

    ssh -oStrictHostKeyChecking=no $BLOCKFREIGHT_SSH_USER@$validator "env validator_name=$validator_name;\
    curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/app.yaml > app.yaml;\
    sed -i -- 's|<VALIDATOR_NAME>|$validator_name|g' app.yaml;\
    sed -i -- 's|'$CONTAINER'|'blockfreight/go-bftx:$1'|g' app.yaml;\
    kubectl apply -f app.yaml && kubectl delete pods --all --grace-period=0 --force;\
    echo old_tag: $CONTAINER;\
    echo new_tag: $1;\
    cat app.yaml;\
    rm app.yaml" <<-'ENDSSH'

ENDSSH
    ((index+=1))
done