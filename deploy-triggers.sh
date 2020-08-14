gcloud alpha builds triggers delete comment-get-trigger --quiet
gcloud alpha builds triggers create github --trigger-config='workflows/trigger-configs/config-get.yaml'