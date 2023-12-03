export $(xargs < .env)

gcloud config set project $GOOGLE_PROJECT_ID

gcloud builds submit --tag $REGION-docker.pkg.dev/$GOOGLE_PROJECT_ID/cloud-run/$CLOUD_RUN_SERVICE:$VERSION_TAG

gcloud run deploy $CLOUD_RUN_SERVICE \
--image $REGION-docker.pkg.dev/$GOOGLE_PROJECT_ID/cloud-run/$CLOUD_RUN_SERVICE:$VERSION_TAG \
--platform=managed \
--allow-unauthenticated \
--region=$REGION \
--project=$GOOGLE_PROJECT_ID

# NOTE: 登入
# gcloud auth login

# NOTE: /r問題，如果sh檔跑過但還是不成功，.env檔也跑一次
# https://blog.csdn.net/Biany0h0/article/details/111160764
# sed -i 's/\r//' deploy.sh

# /r問題方法二
# https://eternalcenter.com/debug-bash-r-command-not-found/
# dos2unix deploy.sh
