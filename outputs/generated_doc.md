# Project Overview
This project consists of various infrastructure configuration files for deploying an application using Kubernetes (Deployment.yml), AWS (main.tf, vpc.tf, test.tf), and Helm (service.yaml). The main goal is to deploy a Siddhi application with the specified configuration on a Kubernetes cluster and provision AWS resources to host the cluster.

## Deployment Configuration Files

### /Users/kishan/Projects/docuforge/test-configs/deployment.yml (Kubernetes)
**Purpose:** This file configures a Kubernetes deployment for a Siddhi application. The deployment will run two replicas of the specified container image, bind to port 8080, and label the pods with `app: siddhi`.

**Key Configuration Options:**
- `replicas`: The number of replicas to run for this deployment. In this case, it's set to 2.
- `selector.matchLabels`: These labels are used by Kubernetes to select and manage pods. In this case, the label `app=siddhi` is used to select all Siddhi application pods.
- `template.spec.containers.image`: The container image for the Siddhi application. In this case, it's set to `siddhi:0.1`.
- `template.spec.containers.ports.containerPort`: The container port that will be exposed and listening for connections. In this case, it's set to 8080.

**Dependencies:** None in the file itself, but it assumes the existence of a Kubernetes cluster with appropriate roles and permissions.

### /Users/kishan/Projects/docuforge/test-configs/main.tf (Terraform)
**Purpose:** This file provisions an AWS `t2.micro` instance using the specified AWS AMIs in the us-east-1 region with the necessary tags.

**Key Configuration Options:**
- `region`: The AWS region where the resources will be created. In this case, it's set to `us-east-1`.
- `ami`: The Amazon Machine Image (AMI) to use for this instance. In this case, it's set to `ami-0c55b159cbfafe1f0`.
- `instance_type`: The type of the AWS EC2 instance. In this case, it's set to `t2.micro`.
- `tags`: Additional metadata for the created resources in AWS. In this case, it sets a tag named `Name` with the value `ExampleInstance`.

**Dependencies:** It assumes the existence of valid AWS credentials and permissions.

### /Users/kishan/Projects/docuforge/test-configs/service.yaml (Helm)
**Purpose:** This file configures a Kubernetes service for the Siddhi application, exposing it on port 80 using the ClusterIP type.

**Key Configuration Options:**
- `apiVersion`: The version of the API being used in this manifest. In this case, it's set to `v1`.
- `kind`: The kind of resource being created (in this case, a service).
- `metadata.name`: The name for the Kubernetes resource. In this case, it's set to `my-app-service`.
- `spec.type`: The type of service to create. In this case, it's set to `ClusterIP`, which means the service is only accessible within the cluster.
- `spec.selector.app`: The label selector for the service, in this case, it selects pods with the label `app=my-app`.
- `spec.ports`: The ports exposed by the service. In this case, it exposes port 80 (TCP) and forwards connections to port 8080 inside the pods.

**Dependencies:** Assumes the existence of a Kubernetes cluster with appropriate roles and permissions.

### /Users/kishan/Projects/docuforge/test-configs/vpc.tf (Terraform)
**Purpose:** This file provisions an AWS VPC with the specified CIDR block and tags in the us-east-1 region.

**Key Configuration Options:**
- `cidr_block`: The CIDR block for the new Amazon Virtual Private Cloud (VPC). In this case, it's set to `10.0.0.0/16`.
- `tags`: Additional metadata for the created resources in AWS. In this case, it sets a tag named `Name` with the value `MainVPC`.

**Dependencies:** It assumes the existence of valid AWS credentials and permissions.

## Architecture Diagram

The following architecture diagram represents the high-level structure of the application and its dependencies:

```mermaid
graph LR
subgraph Kubernetes
  A[Kubernetes Cluster]
    B[Deployment.yml] -->|Configures Siddhi App Deployment| A
end
subgraph AWS
  C[VPC (us-east-1)]
    D[Instance (t2.micro)] -->|Provisioned by main.tf| C
end
subgraph Helm
  E[Helm Chart]
    F[service.yaml] -->|Configures Siddhi App Service| E
end
A -- Deploys --> B
D -- Managed by --> C
F -- Configures --> E
```