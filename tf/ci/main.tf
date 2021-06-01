# Es crea un trigger específic en Google Cloud Build per a que s'executi cada vegada que es creï una Pull Request
# en el repositori, sobre la qual, es passaran els tests unitaris definits en el codi Go. No s'executarà aquest trigger
# quan es detectin canvis en els directoris k8s i tf ja que no contenen codi de l'aplicació.
resource "google_cloudbuild_trigger" "ci_tests_rigger" {
  name          = "tfg-uoc-ci-tests-app"
  description   = "Trigger que s'executarà en cada nova Pull Request creada i commits posteriors"
  filename      = "tf/ci/cloudbuild-tests.yaml"
  ignored_files = ["k8s/**", "tf/**"]

  github {
    owner = "asaldana-uoc"
    name  = "tfg-uoc-app"
    pull_request {
      branch          = ".*"
      comment_control = "COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY"
    }
  }
}

# Aquest un trigger de Google Cloud Build construirà una imatge de Docker cada vegada que hi hagi un commit
# nou en la branca principal main. No s'executarà aquest procés automatitzat quan els canvis hagin succeït
# als directoris k8s i tf ja que no contenen codi de l'aplicació.
resource "google_cloudbuild_trigger" "ci_build_trigger" {
  name          = "tfg-uoc-ci-build-app"
  description   = "Trigger que s'executarà cada nou commit a la branca main per construir la imatge de docker"
  filename      = "tf/ci/cloudbuild-build.yaml"
  ignored_files = ["k8s/**", "tf/**"]

  github {
    owner = "asaldana-uoc"
    name  = "tfg-uoc-app"
    push {
      branch = "main"
    }
  }
}


# Utilitzem el proveïdor google al haver de crear recursos a GCP
provider "google" {
  project = "tfg-uoc-313418"
}

# Aquesta secció es per definir la ubicació on s'emmagatzemarà l'estat de terraform.
# Utilitzem el bucket tfg-uoc-tfstate-eu i desem la configuració de l'aplicación en la ruta tfg-uoc-app/ci
terraform {
  backend "gcs" {
    bucket = "tfg-uoc-tfstate-eu"
    prefix = "tfg-uoc-app/ci"
  }
}
