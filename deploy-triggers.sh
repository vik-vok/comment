CONF_DIR='workflows/trigger-configs/'
REPO_NAME="comment"
REPO_OWNER="vik-vok"
array=(
#  'comment-create-trigger':"${CONF_DIR}create.yaml"
    'comment-get-trigger':"${CONF_DIR}get.yaml"
  #  'comment-get-all-trigger':"${CONF_DIR}get-all.yaml"
  #  'comment-update-trigger':"${CONF_DIR}update.yaml"
)

for i in "${array[@]}"; do
  IFS=":"
  set -- ${i}

  TRIGGER_NAME=${1}
  #  TRIGGER_CONF=${2}
  echo "#### Generating Trigger ${TRIGGER_NAME}"

    gcloud alpha builds triggers delete "${TRIGGER_NAME}" --quiet
  gcloud beta builds triggers create github \
    --repo-name="${REPO_NAME}" \
    --repo-owner="${REPO_OWNER}" \
    --included-files="functions/CommentGet.go" \
    --name="${TRIGGER_NAME}" \
    --branch-pattern="^master$" \
    --build-config 'workflows/cloudbuild.yaml' \
    --substitutions _CLOUD_FUNC_NAME=comment-get,_GO_FUNC_NAME=CommentGet
done
