#!/bin/bash

# isolate previous tag as variable for sed replacement
# read existing app.yaml (which governs the k8s cluster)
OLD_TAG=`curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/app.yaml | \
# grep the blockfreight docker image from the app.yaml file 
# output = image: 'blockfreight/go-bftx:rc1'
grep blockfreight/go-bftx:` | \
# Isolate the docker image + tag
# output = blockfreight/go-bftx:rc1
awk '{print $2}' | \
# Isolate the docker tag
# output = rc1
sed -e 's#.*:\(\)#\1#'`

# travis.yml builds a new docker image. When travis.yml calls this script, it sets pass in the newly built docker image tag and sets it as a variable
# The most recent tag that I had travis build was ci-cd-7b50eb7, which is the branch name, followed by the git commit sha
NEW_TAG=$1

# Index is used to slice the following arrays during the below for loop
index=0
# The top two arrays contain information specific to each bftx node
private_keys=( $PRIVATE_KEY_BFTX0 $PRIVATE_KEY_BFTX1 $PRIVATE_KEY_BFTX2 $PRIVATE_KEY_BFTX3 )
private_node_keys=( $PRIVATE_NODE_KEY_BFTX0 $PRIVATE_NODE_KEY_BFTX1 $PRIVATE_NODE_KEY_BFTX2 $PRIVATE_NODE_KEY_BFTX3 )
# This array contains the k8s master VM ip addresses, which are used by the for loop to ssh into each master, and execute the pipeline
bftx_master_ip_array=( $BFTX0_MASTER_IP $BFTX1_MASTER_IP $BFTX2_MASTER_IP $BFTX3_MASTER_IP )

# loop through ip addresses of each k8s master VM, four total iterations (one for each bftx node)
for validator in  "${bftx_master_ip_array[@]}"
do

    validator_name=bftx${index}
    # slice the arrays for a single value with respect to the corresponding bftx node
    private_key=${private_keys[index]}
    private_node_key=${private_node_keys[index]}

    # ssh into a master k8s VM
    ssh -oStrictHostKeyChecking=no $BLOCKFREIGHT_SSH_USER@$validator "\
    # Set environment variables 
    env validator_name=$validator_name;\
    env private_key=$private_key;\
    env private_node_key=$private_node_key;\
    rm app.yaml;\
    kubectl delete secrets --all --grace-period=0 --force;\
    # Create secrets that tendermint uses to connect the nodes
    kubectl create secret generic node.private.keys --from-literal=privateKey=$private_key --from-literal=privateNodeKey=$private_node_key --from-literal=validatorName=$validator_name;\
    # Download the most up to date app.yaml
    curl https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/app.yaml > app.yaml;\
    # Replace the old docker image tag (rc1) with the newly build docker image tag (ci-cd-7b50eb7)
    # The line will now read image:blockfreight/go-bftx:ci-cd-7b50eb7
    sed -i -- 's|'$OLD_TAG'|'$NEW_TAG'|g' app.yaml;\
    # Apply the changes. k8s will see that there is a new docker image, and bring the k8s pods offline one at a time
    kubectl apply -f app.yaml;" <<-'ENDSSH'


ENDSSH
    # Add one to the index variable so that the proper array slice is returned above
    ((index+=1))
done