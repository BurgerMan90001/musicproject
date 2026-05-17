
# FOR LOCAL DEVELPMENT ONLY
PORT=8081 && ADC=~/.config/gcloud/application_default_credentials.json && \
docker run  \
-p ${PORT}:${PORT} \
-e PORT=${PORT} \
-e GOOGLE_APPLICATION_CREDENTIALS=/tmp/keys/google_credentials.json \
-v ${ADC}:/tmp/keys/google_credentials.json:ro \
us-central1-docker.pkg.dev/high-cistern-488613-c2/cloud-run-source-deploy/songsled-api:latest
