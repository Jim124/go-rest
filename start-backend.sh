BACKEND_IMAGE="go-rest-backend"
BACKEND_CONTAINER_NAME="go-backend"
source .env.db
LOCAL_HOST_PORT=8080
CONTAINER_PORT=8080
source .env.network

if [ "$(docker ps -aq -f name=$BACKEND_CONTAINER_NAME)" ]; then
    echo "A backend container with the name $BACKEND_CONTAINER_NAME already exists."
    echo "Thee container will be removed when stopped."
    echo "To stop the container, run docker kill go-backend"
    exit 1
fi

docker build -t $BACKEND_IMAGE -f Dockerfile.dev . 

docker run --rm -d --name $BACKEND_CONTAINER_NAME \
  --env-file "dev.env" \
  -p $LOCAL_HOST_PORT:$CONTAINER_PORT \
  --network $NETWORK_NAME \
  $BACKEND_IMAGE
