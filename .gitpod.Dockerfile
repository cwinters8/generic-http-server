FROM gitpod/workspace-full:latest

RUN sudo groupadd docker && sudo usermod -aG docker $USER && newgrp docker && \
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.15.0/kind-linux-amd64 && chmod +x ./kind && \
sudo mv ./kind /usr/local/bin/kind && kind create cluster
# sudo apt update -y && sudo apt install snapd -y && sudo snap install microk8s --classic
CMD /bin/bash
