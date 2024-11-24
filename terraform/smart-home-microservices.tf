resource "helm_release" "devices" {
  name      = "devices-service"
  namespace = "default"
  chart     = "../charts/smart-home-microservices/devices"

  depends_on = [helm_release.kafka]

  values = [
    file("../charts/smart-home-microservices/devices/values.yaml"),
  ]
}

resource "helm_release" "telemetry" {
  name      = "telemetry-service"
  namespace = "default"
  chart     = "../charts/smart-home-microservices/telemetry"

  depends_on = [helm_release.kafka]

  values = [
    file("../charts/smart-home-microservices/telemetry/values.yaml"),
  ]
}
