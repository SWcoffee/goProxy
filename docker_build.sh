docker build -t goproxy .
# Step 2: Tag the Docker image
docker tag goproxy:latest swcoffee/goproxy:latest

 Step 4: Push the Docker image to Docker Hub
 docker push swcoffee/goproxy:latest