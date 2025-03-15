source .env.volume
source .env.network

if [ "$(docker volume ls -q -f name=$VOLUME_NAME)" ]; then
    echo "A volume with the name $VOLUME_NAME already existed, skipping this step."
else
    docker volume create $VOLUME_NAME
fi

if [ "$(docker network ls -q -f name=$NETWORK_NAME)" ]; then
    echo "A network with the name $NETWORK_NAME already existed, skipping this step."
else
    docker network create $NETWORK_NAME
fi