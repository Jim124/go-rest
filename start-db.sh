MYSQL_IMAGE_NAME="mysql"
MYSQL_TAG="8.0.39"

source .env.db

#root credentials
ROOT_PASSWORD="root123"
#connectivity
source .env.network
LOCAL_HOST_PORT=3306
CONTAINER_PORT=3306
#storage
source .env.volume
VOLUME_CONTAINER_PATH="/data/mysql"

source setup.sh

if [ "$(docker ps -aq -f name=$DB_CONTAINER_NAME)" ]; then
    echo "The container with the name $DB_CONTAINER_NAME already existed."
    echo "The container will be removed when stopped."
    echo "To stop the container, run docker kill mysql8"
    exit 1
fi

docker run --rm -d --name $DB_CONTAINER_NAME \
  -e MYSQL_ROOT_PASSWORD=$ROOT_PASSWORD \
  -p $LOCAL_HOST_PORT:$CONTAINER_PORT \
  -v $VOLUME_NAME:$VOLUME_CONTAINER_PATH \
  --network $NETWORK_NAME \
  $MYSQL_IMAGE_NAME:$MYSQL_TAG


