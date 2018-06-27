#!/bin/bash

# isolate previous tag as variable for sed replacement
PREV_TAG=`curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/app.yaml | grep blockfreight/go-bftx: | awk '{print $2}' | sed -e 's#.*:\(\)#\1#'`

index=0
for validator in $BFTX0_MASTER_IP $BFTX1_MASTER_IP $BFTX2_MASTER_IP $BFTX3_MASTER_IP
do
    validator_name=bftx${index}

    ssh -oStrictHostKeyChecking=no $BLOCKFREIGHT_SSH_USER@$validator "env validator_name=$validator_name;\
    curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/app.yaml > app.yaml;\
    sed -i -- 's|<VALIDATOR_NAME>|$validator_name|g' app.yaml;\
    sed -i -- 's|'$PREV_TAG'|'$1'|g' app.yaml;\
    echo old_tag: $PREV_TAG
    echo new_tag: $1
    cat app.yaml | grep $1
    kubectl apply -f app.yaml && kubectl delete pods --all --grace-period=0 --force;\
    rm app.yaml" <<-'ENDSSH'

ENDSSH
    ((index+=1))
done