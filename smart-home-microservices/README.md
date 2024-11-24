## Запуск

```bash
docker build -t devices -f build/devices/Dockerfile .
docker build -t telemetry -f build/telemetry/Dockerfile .
```

### Загрузка образов в minicube

```bash
minikube image load telemetry:latest
minikube image load devices:latest
```

### Обновление зависимостей helm

```bash
helm dependency update ../charts/smart-home-microservices/devices
helm dependency update ../charts/smart-home-microservices/telemetry
helm dependency update ../charts/smart-home-monolith
```

### Запуск через heml

```bash
helm install devices ../charts/smart-home-microservices/devices
helm install telemetry ../charts/smart-home-microservices/devices
```

### Запуск через terraform

```bash
cd ../terraform
terraform init
terraform apply
```
