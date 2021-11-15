This project shows an example go application that uses skaffold for local development.

Install skaffold:

    https://skaffold.dev/docs/install/
    
Configure your local environment to use the minikube docker daemon:

    eval $(minikube docker-env)
    
Initialize Skaffold:
    skaffold init

Create namespace:

    kubectl create namespace feature

To build/deploy this app, from the project root run:

    skaffold dev

Once deployed, hit the endpoint with curl/postman using kong URL:

    curl -i <your kong url>/hello
    
If you now make a change to the code, skaffold will detect the change when the file is saved and rebuild/redeploy the 
app automatically.

Skaffold documentation:  https://skaffold.dev/docs/