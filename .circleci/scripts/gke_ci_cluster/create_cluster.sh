export GCP_PROJECT_NAME=determined-ai
export GKE_CLUSTER_NAME=gke-circleci-newnet
export GKE_REGION=us-west1
export GKE_MACHINE_TYPE=n1-standard-8
export GKE_NUM_NODES=1
export GCS_NETWORK_NAME=gke-circleci-vpc-network

export GKE_GPU_NODE_POOL_NAME=gke-circleci-compute-newnet

export BASTION_INSTANCE_NAME=gke-circleci-bastion-newnet
export BASTION_INSTANCE_ZONE="$GKE_REGION-b"
export BASTION_INSTANCE_TYPE=e2-micro

gcloud container clusters create "$GKE_CLUSTER_NAME" \
        --project "$GCP_PROJECT_NAME" \
        --region "$GKE_REGION" \
        --network="$GCS_NETWORK_NAME" \
        --enable-autoscaling \
        --enable-autoprovisioning \
        --max-cpu=7 \
        --max-memory=25 \
        --num-nodes="$GKE_NUM_NODES" \
        --enable-master-authorized-networks \
        --enable-ip-alias \
        --enable-private-nodes \
        --enable-private-endpoint \
        --master-ipv4-cidr 172.16.0.0/28 \
        --cluster-version=latest \
        --machine-type="$GKE_MACHINE_TYPE" \
        --maintenance-window 8:00

gcloud container node-pools create "$GKE_GPU_NODE_POOL_NAME" \
        --cluster "$GKE_CLUSTER_NAME" \
        --project "$GCP_PROJECT_NAME" \
        --region="$GKE_REGION" \
        --machine-type="$GKE_COMPUTE_MACHINE_TYPE" \
        --scopes=storage-full \
        --enable-autoscaling \
        --enable-autoprovisioning \
        --enable-autorepair

gcloud compute instances create "$BASTION_INSTANCE_NAME" \
        --project "$GCP_PROJECT_NAME" \
        --zone="$BASTION_INSTANCE_ZONE" \
        --machine-type="$BASTION_INSTANCE_TYPE" \
        --network-interface=no-address,subnet="$GCS_NETWORK_NAME"
