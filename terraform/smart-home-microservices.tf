resource "helm_release" "devices" {
  name      = "devices-service"
  namespace = "default"
  chart     = "../charts/smart-home-microservices/devices"
}

resource "helm_release" "telemetry" {
  name      = "telemetry-service"
  namespace = "default"
  chart     = "../charts/smart-home-microservices/telemetry"
}
