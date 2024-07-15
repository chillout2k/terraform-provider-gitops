<<<<<<< HEAD
# terraform-provider-gitops
=======
# gitops Terraform Provider für gitops-REST-API

## Vorbereitungen
Lokale `~/.terraformrc` anlegen

Mit dieser Terraform-Konfiguration wird ein Override erzeugt, damit der Terraform die Provider-Binary nicht im Internet sucht sondern lokal auf Platte:
```
provider_installation {

  dev_overrides {
      "hashicorp.com/edu/gitops" = "/Users/dominik.chilla/go/bin"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

## Provider kompilieren
```
> go install
```
Binary `terraform-provider-gitops` sollte unter `~/go/bin/` zu finden sein

## Starte die gitops-REST-API
siehe Doku unter ../server/

## Provider in action
### Erstelle das Terraform-Skript `main.tf`:
```
terraform {
  required_providers {
    gitops = {
      source = "hashicorp.com/edu/gitops"
    }
  }
}

provider "gitops" {
  host = "http://localhost:8000"
}

resource "gitops_instance" "test1" {
  instance_name = "terraform provisioned test1"
  orderer_id    = "dein.email@adresse"
  bits_account  = 1234
  service_id    = 4321
  replica_count = 4
  version       = "3.2.*"
  some_value    = "test instance 1"
}

output "instance1" {
  value = gitops_instance.test1
}
```

### Terraform plan
```
> terraform plan
╷
│ Warning: Provider development overrides are in effect
│ 
│ The following provider development overrides are set in the CLI configuration:
│  - hashicorp.com/edu/gitops in /Users/dominik.chilla/go/bin
│ 
│ The behavior may therefore not match any released version of the provider and applying changes may cause the
│ state to become incompatible with published releases.
╵

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated
with the following symbols:
  + create

Terraform will perform the following actions:

  # gitops_instance.test1 will be created
  + resource "gitops_instance" "test1" {
      + bits_account  = 1234
      + instance_id   = (known after apply)
      + instance_name = "terraform provisioned test1"
      + last_updated  = (known after apply)
      + order_time    = (known after apply)
      + orderer_id    = "dein.email@adresse"
      + replica_count = 4
      + service_id    = 4321
      + some_value    = "test instance 1"
      + stage         = (known after apply)
      + version       = "3.2.*"
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Changes to Outputs:
  + instance1 = {
      + bits_account  = 1234
      + instance_id   = (known after apply)
      + instance_name = "terraform provisioned test1"
      + last_updated  = (known after apply)
      + order_time    = (known after apply)
      + orderer_id    = "dein.email@adresse"
      + replica_count = 4
      + service_id    = 4321
      + some_value    = "test instance 1"
      + stage         = (known after apply)
      + version       = "3.2.*"
    }

───────────────────────────────────────────────────────────────────────────────────────────────────────────────

Note: You didn't use the -out option to save this plan, so Terraform can't guarantee to take exactly these
actions if you run "terraform apply" now.
```

### Terraform apply
```
> terraform apply
╷
│ Warning: Provider development overrides are in effect
│ 
│ The following provider development overrides are set in the CLI configuration:
│  - hashicorp.com/edu/gitops in /Users/dominik.chilla/go/bin
│ 
│ The behavior may therefore not match any released version of the provider and applying changes may cause the
│ state to become incompatible with published releases.
╵

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated
with the following symbols:
  + create

Terraform will perform the following actions:

  # gitops_instance.test1 will be created
  + resource "gitops_instance" "test1" {
      + bits_account  = 1234
      + instance_id   = (known after apply)
      + instance_name = "terraform provisioned test1"
      + last_updated  = (known after apply)
      + order_time    = (known after apply)
      + orderer_id    = "dein.email@adresse"
      + replica_count = 4
      + service_id    = 4321
      + some_value    = "test instance 1"
      + stage         = (known after apply)
      + version       = "3.2.*"
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Changes to Outputs:
  + instance1 = {
      + bits_account  = 1234
      + instance_id   = (known after apply)
      + instance_name = "terraform provisioned test1"
      + last_updated  = (known after apply)
      + order_time    = (known after apply)
      + orderer_id    = "dein.email@adresse"
      + replica_count = 4
      + service_id    = 4321
      + some_value    = "test instance 1"
      + stage         = (known after apply)
      + version       = "3.2.*"
    }

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

gitops_instance.test1: Creating...
gitops_instance.test1: Creation complete after 1s

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.

Outputs:

instance1 = {
  "bits_account" = 1234
  "instance_id" = "e449ed1c-820b-43c7-b1ab-98caee3e317f"
  "instance_name" = "terraform provisioned test1"
  "last_updated" = "Friday, 05-Jul-24 22:29:31 CEST"
  "order_time" = "Fri Jul  5 22:29:30 2024"
  "orderer_id" = "dein.email@adresse"
  "replica_count" = 4
  "service_id" = 4321
  "some_value" = "test instance 1"
  "stage" = "prod"
  "version" = "3.2.*"
}
```
**Die erzeugte gitops-Instanz hat die ID: `e449ed1c-820b-43c7-b1ab-98caee3e317f`**

Logs der gitops-REST-API:
```
INFO:     127.0.0.1:49740 - "POST /instances HTTP/1.1" 200 OK
```

### Plan ändern und updaten
In der `main.tf` einen Wert von `resource "gitops_instance" "test1"` ändern, z.B. `replica_count` auf 6, speichern und `terraform plan` ausführen:
```
> terraform plan 
╷
│ Warning: Provider development overrides are in effect
│ 
│ The following provider development overrides are set in the CLI configuration:
│  - hashicorp.com/edu/gitops in /Users/dominik.chilla/go/bin
│ 
│ The behavior may therefore not match any released version of the provider and applying changes may cause the
│ state to become incompatible with published releases.
╵
gitops_instance.test1: Refreshing state...

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated
with the following symbols:
  ~ update in-place

Terraform will perform the following actions:

  # gitops_instance.test1 will be updated in-place
  ~ resource "gitops_instance" "test1" {
      ~ last_updated  = "Friday, 05-Jul-24 22:29:31 CEST" -> (known after apply)
      ~ replica_count = 4 -> 6
        # (9 unchanged attributes hidden)
    }

Plan: 0 to add, 1 to change, 0 to destroy.

Changes to Outputs:
  ~ instance1 = {
      ~ last_updated  = "Friday, 05-Jul-24 22:29:31 CEST" -> (known after apply)
      ~ replica_count = 4 -> 6
        # (9 unchanged attributes hidden)
    }

───────────────────────────────────────────────────────────────────────────────────────────────────────────────

Note: You didn't use the -out option to save this plan, so Terraform can't guarantee to take exactly these
actions if you run "terraform apply" now.
```
Siehst Du die vorzunehmenden Änderungen? Anschließend `terraform apply` ausführen:
```
> terraform apply
╷
│ Warning: Provider development overrides are in effect
│ 
│ The following provider development overrides are set in the CLI configuration:
│  - hashicorp.com/edu/gitops in /Users/dominik.chilla/go/bin
│ 
│ The behavior may therefore not match any released version of the provider and applying changes may cause the
│ state to become incompatible with published releases.
╵
gitops_instance.test1: Refreshing state...

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated
with the following symbols:
  ~ update in-place

Terraform will perform the following actions:

  # gitops_instance.test1 will be updated in-place
  ~ resource "gitops_instance" "test1" {
      ~ last_updated  = "Friday, 05-Jul-24 22:29:31 CEST" -> (known after apply)
      ~ replica_count = 4 -> 6
        # (9 unchanged attributes hidden)
    }

Plan: 0 to add, 1 to change, 0 to destroy.

Changes to Outputs:
  ~ instance1 = {
      ~ last_updated  = "Friday, 05-Jul-24 22:29:31 CEST" -> (known after apply)
      ~ replica_count = 4 -> 6
        # (9 unchanged attributes hidden)
    }

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

gitops_instance.test1: Modifying...
gitops_instance.test1: Modifications complete after 0s

Apply complete! Resources: 0 added, 1 changed, 0 destroyed.

Outputs:

instance1 = {
  "bits_account" = 1234
  "instance_id" = "e449ed1c-820b-43c7-b1ab-98caee3e317f"
  "instance_name" = "terraform provisioned test1"
  "last_updated" = "Friday, 05-Jul-24 22:35:49 CEST"
  "order_time" = "Fri Jul  5 22:29:30 2024"
  "orderer_id" = "dein.email@adresse"
  "replica_count" = 6
  "service_id" = 4321
  "some_value" = "test instance 1"
  "stage" = "prod"
  "version" = "3.2.*"
}
```
Siehe Logs der gitops-REST-API:
```
INFO:     127.0.0.1:49817 - "GET /instances/e449ed1c-820b-43c7-b1ab-98caee3e317f HTTP/1.1" 200 OK
INFO:     127.0.0.1:49820 - "PUT /instances/e449ed1c-820b-43c7-b1ab-98caee3e317f HTTP/1.1" 200 OK
```

### Terraform destroy
gitops Instanz wieder einreißen:
```
> terraform destroy
╷
│ Warning: Provider development overrides are in effect
│ 
│ The following provider development overrides are set in the CLI configuration:
│  - hashicorp.com/edu/gitops in /Users/dominik.chilla/go/bin
│ 
│ The behavior may therefore not match any released version of the provider and applying changes may cause the
│ state to become incompatible with published releases.
╵
gitops_instance.test1: Refreshing state...

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated
with the following symbols:
  - destroy

Terraform will perform the following actions:

  # gitops_instance.test1 will be destroyed
  - resource "gitops_instance" "test1" {
      - bits_account  = 1234 -> null
      - instance_id   = "e449ed1c-820b-43c7-b1ab-98caee3e317f" -> null
      - instance_name = "terraform provisioned test1" -> null
      - last_updated  = "Friday, 05-Jul-24 22:35:49 CEST" -> null
      - order_time    = "Fri Jul  5 22:29:30 2024" -> null
      - orderer_id    = "dein.email@adresse" -> null
      - replica_count = 6 -> null
      - service_id    = 4321 -> null
      - some_value    = "test instance 1" -> null
      - stage         = "prod" -> null
      - version       = "3.2.*" -> null
    }

Plan: 0 to add, 0 to change, 1 to destroy.

Changes to Outputs:
  - instance1 = {
      - bits_account  = 1234
      - instance_id   = "e449ed1c-820b-43c7-b1ab-98caee3e317f"
      - instance_name = "terraform provisioned test1"
      - last_updated  = "Friday, 05-Jul-24 22:35:49 CEST"
      - order_time    = "Fri Jul  5 22:29:30 2024"
      - orderer_id    = "dein.email@adresse"
      - replica_count = 6
      - service_id    = 4321
      - some_value    = "test instance 1"
      - stage         = "prod"
      - version       = "3.2.*"
    } -> null

Do you really want to destroy all resources?
  Terraform will destroy all your managed infrastructure, as shown above.
  There is no undo. Only 'yes' will be accepted to confirm.

  Enter a value: yes

gitops_instance.test1: Destroying...
gitops_instance.test1: Destruction complete after 0s

Destroy complete! Resources: 1 destroyed.
```
Siehe Logs der gitops-REST-API:
```
INFO:     127.0.0.1:49832 - "GET /instances/e449ed1c-820b-43c7-b1ab-98caee3e317f HTTP/1.1" 200 OK
INFO:     127.0.0.1:49834 - "DELETE /instances/e449ed1c-820b-43c7-b1ab-98caee3e317f HTTP/1.1" 200 OK
```
>>>>>>> 83869bd (init)
