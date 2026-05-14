# Wisecow Application - Containerization & K8s Deployment

This repository contains the solution for the Wisecow DevOps Trainee assessment. It demonstrates the containerization of a bash-based application, Kubernetes deployment manifests, and an automated CI/CD pipeline.

## Project Overview

The objective of this project is to containerize the `wisecow.sh` application, create Kubernetes manifests for deployment, and automate the build and deployment process using GitHub Actions.

### Technologies Used
* **Docker:** Containerization of the application.
* **Kubernetes:** Manifests for Deployment, Service, and Ingress (TLS).
* **GitHub Actions:** CI/CD pipeline automation.
* **Linux utilities:** `fortune`, `cowsay`, `netcat`.

---

## 1. Dockerization
The application is a shell script (`wisecow.sh`) that serves ASCII art over a specific port. 
* A lightweight `ubuntu:22.04` base image was used.
* Required dependencies (`fortune-mod`, `cowsay`, `netcat-openbsd`) are installed via the Dockerfile.
* The image exposes port `4499`.

## 2. Kubernetes Deployment
The `k8s/` directory contains the necessary manifests to deploy the application to a Kubernetes cluster (like Minikube or Kind).

* **`deployment.yaml`**: Deploys the Docker image with 2 replicas for high availability.
* **`service.yaml`**: Exposes the deployment internally via a ClusterIP service on port 80, routing to target port 4499.
* **`ingress.yaml` (TLS Challenge Goal)**: Secures the application using an NGINX Ingress controller. It configures TLS termination using a Kubernetes secret (`wisecow-tls`) and routes traffic for `wisecow.local`.

### *Note on TLS Implementation:*
To fully enable TLS locally, you must generate a secret and apply it to your cluster:
```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=wisecow.local"
kubectl create secret tls wisecow-tls --key tls.key --cert tls.crt
