# 1 stop and remove container
source .env.db
#2 remove volume
source .env.volume
#3 remove network
source .env.network

if [ "$(docker ps -aq -f name=$DB_CONTAINER_NAME)" ]; then
    echo "removing container $DB_CONTAINER_NAME"
    docker kill $DB_CONTAINER_NAME
else
    echo "the container with the name $DB_CONTAINER_NAME does not exist, skipping this step."
fi

if [ "$(docker volume ls -f name=$VOLUME_NAME)" ]; then
    echo "Removing volume $VOLUME_NAME"
    docker volume rm $VOLUME_NAME
else
    echo "the volume with the name $VOLUME_NAME does not exist, skipping this step."
fi

if [ "$(docker network ls -f name=$NETWORK_NAME)" ]; then
    echo "Removing network $NETWORK_NAME"
    docker network rm $NETWORK_NAME
else
    echo "the network with the name $NETWORK_NAME does not exist, skipping this step."
fi
