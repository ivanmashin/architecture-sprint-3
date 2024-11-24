resource "helm_release" "postgresql" {
  name       = "postgresql"
  namespace  = "default"
  chart      = "postgresql"
  repository = "oci://registry-1.docker.io/bitnamicharts"
  version    = "15.5.20"
}
