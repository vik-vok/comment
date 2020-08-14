CONF_DIR='workflows/trigger-configs/'

array=(
  'comment-create-trigger':"${CONF_DIR}create.yaml"
  'comment-get-trigger':"${CONF_DIR}get.yaml"
  'comment-get-all-trigger':"${CONF_DIR}get-all.yaml"
  'comment-update-trigger':"${CONF_DIR}update.yaml"
)

for i in "${array[@]}"; do IFS=":";
  set -- ${i};

  TRIGGER_NAME=${1}
  TRIGGER_CONF=${2}
  echo "#### Generating Trigger ${TRIGGER_NAME}"

  gcloud alpha builds triggers delete "${TRIGGER_NAME}" --quiet
  gcloud alpha builds triggers create github --trigger-config="${TRIGGER_CONF}"
done
