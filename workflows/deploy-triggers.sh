BUILD_CONF='cloudbuild_template.yaml'
REPO_NAME="comment"
REPO_OWNER="vik-vok"

array=(
  'comment-create':'CommentCreate.go'
  'comment-get':'CommentGet.go'
  'comment-get-all':'CommentGetAll.go'
  'comment-get-update':'CommentUpdate.go'
)

for i in "${array[@]}"; do
  IFS=":"
  set -- ${i}

  CLOUD_FUNC_NAME=${1}
  GO_FUNC_NAME=${2}
  TRIGGER_NAME="${CLOUD_FUNC_NAME}-trigger"
  echo "#### Generating Trigger ${TRIGGER_NAME}"

  gcloud alpha builds triggers delete "${TRIGGER_NAME}" --quiet
  gcloud beta builds triggers create github \
    --repo-name="${REPO_NAME}" \
    --repo-owner="${REPO_OWNER}" \
    --included-files="../functions/${GO_FUNC_NAME}" \
    --name="${TRIGGER_NAME}" \
    --branch-pattern="^master$" \
    --build-config=${BUILD_CONF} \
    --substitutions _CLOUD_FUNC_NAME="${CLOUD_FUNC_NAME}",_GO_FUNC_NAME="${GO_FUNC_NAME}"
done
