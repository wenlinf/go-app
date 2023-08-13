# go-app
## Overview
This is a Go application that interacts with the microplate reader /get_measurement api. The /get_measurement api returns a matrix of 8*12. The app calls the service every minute and returns the average of the values in the matrix and prints it out. 

## How to run the app
#### Step 1
Follow this guide to get Shifu up and running: https://shifu.dev/docs/tutorials/demo-install/
#### Step 2
Follow this guide to get the digital twin of the microplate reader up and running: https://shifu.dev/docs/tutorials/demo-try/#3-interact-with-the-microplate-reader
#### Step 3
```
cd go-app
sudo kubectl apply -f deployment.yaml
``` 
To see the logs:
Run the following command to get the name of the pod
``` 
sudo kubectl get pods 
``` 
Then run the following command with the pod name you get from the previous step
```
kubectl logs <pod-name>
```
