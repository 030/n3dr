#!/bin/bash -e

docker() {
    echo "Installing docker..."
    apt-get update
    apt-get install -y \
        apt-transport-https \
        ca-certificates \
        curl \
        gnupg \
        lsb-release
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor --yes -o /usr/share/keyrings/docker-archive-keyring.gpg
    echo \
        "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
        $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list >/dev/null
    apt-get update
    apt-get install -y docker-ce docker-ce-cli containerd.io
    sudo systemctl restart docker
    sudo docker run hello-world
    sudo usermod -aG docker vagrant
    whoami
}

n3dr() {
    cd /vagrant
    sed -i "s|^\(trap clean EXIT\)|#\1|" test/integration-tests.sh
    sed -i "s|  upload|echo \$PASSWORD \&\& exit 0 #upload|" test/integration-tests.sh
    sed -i "s|  build|#build|" test/integration-tests.sh
    ./test/integration-tests.sh
    cd -
    nexus3ip=$(ip a | grep 192 | awk '{ print $2 }' | sed -e "s|\/24||g")
    echo "Navigate to ${nexus3ip}:9999 and login with 'admin' and the provided password"
}

main() {
    docker
    n3dr
}

main
