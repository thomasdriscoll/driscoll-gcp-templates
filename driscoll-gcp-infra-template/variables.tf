// GKE Cluster variables
variable "gke_username" {
  default     = ""
  description = "gke username"
}

variable "gke_password" {
  default     = ""
  description = "gke password"
}

variable "gke_num_nodes" {
  default     = 2
  description = "number of gke nodes"
}

// VPC Variables
variable "project_id" {
  description = "project ID"
}

variable "region" {
  description = "region"
}

