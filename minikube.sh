minikube start
minikube addons enable ingress

MINIKUBE_IP=$(minikube ip)
SSH_KEY="~/.minikube/machines/minikube/id_rsa"

kubectl config view --flatten > /tmp/kubeconfig
ssh -i $SSH_KEY docker@$MINIKUBE_IP "mkdir ~/.kube"
minikube cp /tmp/kubeconfig /home/docker/.kube/config

cd ./playbook
ansible-playbook install.yaml -e "ansible_host=$MINIKUBE_IP ansible_user=docker ansible_ssh_private_key_file=$SSH_KEY"
