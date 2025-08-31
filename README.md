# DevOps

docker build -t kube-service:local ./apps/service

docker run -d -p 8080:8080 --name kube-service kube-service:local

curl http://localhost:8080
curl http://localhost:8080/metrics

helm install service ./deploy/helm/service
kubectl get pods
kubectl port-forward svc/service 8080:8080

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm install monitoring prometheus-community/kube-prometheus-stack -n monitoring --create-namespace

kubectl port-forward svc/monitoring-grafana 3000:80 -n monitoring
