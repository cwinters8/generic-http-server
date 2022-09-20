FROM gitpod/workspace-full:latest

RUN sudo groupadd docker && sudo usermod -aG docker $USER && newgrp docker && \
sudo apt update -y && sudo apt install snapd -y && sudo snap install microk8s --classic
CMD /bin/bash
