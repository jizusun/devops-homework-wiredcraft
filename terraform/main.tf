# https://cloud.google.com/free

terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "3.5.0"
    }
  }
}

provider "google" {
  credentials = file("wiredcraft-test-devops-4851a96627b4.gcp.json")
  project = "wiredcraft-test-devops"
  region  = "us-central1"
  zone    = "us-central1-c"
}

resource "google_compute_instance" "default" {

  name         = "hugo-vm"
  machine_type = "f1-micro"
  tags         = ["web", "hugo"]

  provisioner "local-exec" {
    command = "echo ${google_compute_instance.default.name}:  ${google_compute_instance.default.network_interface[0].access_config[0].nat_ip} >> ip_address.txt"
  }

  metadata = {
   ssh-keys = "hugo:${file("~/.ssh/wiredcraft.pub")}"
  }

  boot_disk {
    initialize_params {
        image = "centos-cloud/centos-7"
    }
  }

  network_interface {
    network = "default"
    access_config {
    }
  }
}

output "ip" {
    value = google_compute_instance.default.network_interface.0.access_config.0.nat_ip
  }


resource "google_compute_firewall" "default" {
 name    = "hugo-firewall"
 network = "default"

 allow {
   protocol = "tcp"
   ports    = ["80"]
 }
}