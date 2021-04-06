# Golang Webservice Example

# 0. Run Webservice
```
go build -o app  
./app
visit your browser and input `http://localhost:8000/hello`
```

## 1. Vagrant up

```
vagrant init centos/7
vagrant up
```

## 2. Update vm and install podman 

```
sudo yum update
sudo yum -y install podman
```

## 3. Build image
```
sudo podman build . -t quay.io/${YOUR_QUAY_ACCOUNT}/${YOUR_REPO_NAME}:${IMAGE_TAG}
```

## 4. Push it to Quay registry
```
sudo podman login quay.io
sudo podman push quay.io/${YOUR_QUAY_ACCOUNT}/${YOUR_REPO_NAME}:${IMAGE_TAG}
```