## Запуск

```bash
docker build -t devices -f build/devices/Dockerfile .
docker build -t telemetry -f build/telemetry/Dockerfile .
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
