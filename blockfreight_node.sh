if [ -x "$(command -v docker)" ]; then
    echo 'Downloading Blockfreight Go-bftx docker image'
    docker pull blockfreight/go-bftx:rc1
    
    echo 'Initializing Blockfreight Node'
    docker run --entrypoint=bftx -p 46657:46657 -p 8080:8080 blockfreight/go-bftx:rc1 node start
    # command
else
    echo "Go-bftx requires Docker but it appears not to be installed. Please check your system has Docker installed (https://docs.docker.com/install/) before installing Go-bftx. Aborting."
    exit 1;
fi