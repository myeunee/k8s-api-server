# k8s-api-server
![Terraform](https://img.shields.io/badge/Terraform-purple)
![Docker](https://img.shields.io/badge/Docker-2496ED)
![Kubernetes](https://img.shields.io/badge/Kubernetes-darkgreen)
![Golang](https://img.shields.io/badge/Golang-blue)
![MySQL](https://img.shields.io/badge/MySQL-grey)
![GCP](https://img.shields.io/badge/GCP-orange)


Go로 개발된 API 서버가 Terraform을 통해 정의된 인프라를 기반으로 Kubernetes 클러스터와 상호작용하며, GCP에서 리소스를 관리한다. 



### 1️⃣ **SA Key Used By Terraform**
- Terraform은 GCP의 서비스 계정 키를 사용하여 인증을 처리
- 이 서비스 계정은 부여된 권한을 통 GCP API에 접근하고 Terraform을 통해 인프라를 생성한다.

### 2️⃣ **Terraform Apply**
- terraform apply 명령어를 통해 GCP 인프라 리소스를 배포
- Terraform은 GCP API와 상호작용하여 Kubernetes 클러스터, VPC, Cloud Storage 등을 프로비저닝

### 3️⃣ **Authorized API Calls**

### 4️⃣ **Go API Server Access to Kubernetes Cluster**
- Golang API 서버는 Kubernetes 클러스터의 리소스에 접근하여 필요한 서비스를 호출하거나 데이터를 처리
