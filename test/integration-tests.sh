#!/bin/bash -ex

if [ -z "${N3DR_APT_GPG_SECRET}" ]; then
  echo "N3DR_APT_GPG_SECRET should not be empty"
  echo "Create one by running:"
  echo "docker run -v /tmp/gpg-output:/root/.gnupg -v ${PWD}/test/gpg/:/tmp/ --rm -it vladgh/gpg --batch --generate-key /tmp/generate"
  echo "docker run --rm -it -v /tmp/gpg-output:/root/.gnupg -v ${PWD}/test/gpg/:/tmp/ vladgh/gpg --output /tmp/my_rsa_key --armor --export-secret-key joe@foo.bar"
  echo "Enter 'abc' as a password, if the prompt appears"
  printf "export N3DR_APT_GPG_SECRET=\$(sudo cat test/gpg/my_rsa_key | tr '\\\n' ' ' | sed -r \"s|-----[A-Z]+ PGP PRIVATE KEY BLOCK-----||g;s| |\\\\\\\\\\\n|g;s|(.*)|-----BEGIN PGP PRIVATE KEY BLOCK-----\\\1-----END PGP PRIVATE KEY BLOCK-----|g\")"
  echo
  echo "sudo rm -r /tmp/gpg-output"
  echo "rm test/gpg/my_rsa_key"
  echo
  printf "Note: Spaces and enters have to be escaped, i.e. '\\\n'->'\\\\\\\n' and ' '->'\ ' if the token is used in travis."
  exit 1
fi

if [ -z "${NEXUS_VERSION}" ]; then
  echo "NEXUS_VERSION empty, setting it to the default value"
  NEXUS_VERSION=3.44.0
fi

if [ -z "${NEXUS_API_VERSION}" ]; then
  echo "NEXUS_API_VERSION empty, setting it to the default value"
  NEXUS_API_VERSION=v1
fi

if [ -z "${N3DR_DELIVERABLE}" ]; then
  echo "N3DR_DELIVERABLE empty, setting it to the default value"
  N3DR_DELIVERABLE=./n3dr
fi

if [ -z "${N3DR_CLEAN_IN_CASE_OF_SUCCESS_OR_FAILURE}" ]; then
  echo "N3DR_CLEAN_IN_CASE_OF_SUCCESS_OR_FAILURE empty, setting it to the default value"
  N3DR_CLEAN_IN_CASE_OF_SUCCESS_OR_FAILURE=true
fi

readonly DOCKER_REGISTRY_HTTP_CONNECTOR_A=8888
readonly DOCKER_REGISTRY_HTTP_CONNECTOR_B=8887
readonly DOCKER_REGISTRY_HTTP_CONNECTOR_C=8886
readonly DOCKER_REGISTRY_HTTP_CONNECTOR_INTERNAL=8083
readonly DOCKER_URL=http://localhost
readonly DOWNLOAD_LOCATION=/tmp/n3dr
readonly DOWNLOAD_LOCATION_PASS=${DOWNLOAD_LOCATION}-pass
readonly DOWNLOAD_LOCATION_RPROXY=${DOWNLOAD_LOCATION}-rproxy
readonly DOWNLOAD_LOCATION_SYNC=${DOWNLOAD_LOCATION}-sync
readonly DOWNLOAD_LOCATION_SYNC_A=${DOWNLOAD_LOCATION_SYNC}-a
readonly DOWNLOAD_LOCATION_SYNC_B=${DOWNLOAD_LOCATION_SYNC}-b
readonly DOWNLOAD_LOCATION_SYNC_C=${DOWNLOAD_LOCATION_SYNC}-c
readonly HOSTED_REPO_DOCKER=REPO_NAME_HOSTED_DOCKER
readonly HOSTED_REPO_YUM=REPO_NAME_HOSTED_YUM
readonly PORT_NEXUS_A=9999
readonly PORT_NEXUS_B=9998
readonly PORT_NEXUS_C=9997
readonly URL_NEXUS_A=http://localhost:${PORT_NEXUS_A}
readonly URL_NEXUS_B=http://localhost:${PORT_NEXUS_B}
readonly URL_NEXUS_C=http://localhost:${PORT_NEXUS_C}
readonly URL_NEXUS_A_V2="${URL_NEXUS_A/http:\/\//}"

validate() {
  if [ -z "${N3DR_DELIVERABLE}" ]; then
    echo "No deliverable defined. Assuming that 'go run main.go' should be run."
    N3DR_DELIVERABLE="go run main.go"
  fi
  if [ -z "${NEXUS_VERSION}" ] || [ -z "${NEXUS_API_VERSION}" ]; then
    echo "NEXUS_VERSION and NEXUS_API_VERSION should be specified."
    exit 1
  fi
  if [ -d "${DOWNLOAD_LOCATION}" ]; then
    echo "Ensure that ${DOWNLOAD_LOCATION} does not exist"
    exit 1
  fi
}

build() {
  # shellcheck disable=SC1091
  source ./scripts/build.sh
  cd cmd/n3dr
}

startNexus() {
  # shellcheck disable=SC1091
  # as nexus-docker.sh is retrieved remotely
  source ./start.sh "${NEXUS_VERSION}" "${NEXUS_API_VERSION}" "nexus-${1}" "${2}" "${DOWNLOAD_LOCATION_PASS}" "${3}" "${DOCKER_REGISTRY_HTTP_CONNECTOR_INTERNAL}" &>/dev/null
}

nexus() {
  curl -sL https://gist.githubusercontent.com/030/666c99d8fc86e9f1cc0ad216e0190574/raw/6dd6ce267cba17139d4da0ece908b8c67c14bd21/nexus-docker.sh -o start.sh
  chmod +x start.sh

  startNexus a ${PORT_NEXUS_A} ${DOCKER_REGISTRY_HTTP_CONNECTOR_A} &
  startNexus b ${PORT_NEXUS_B} ${DOCKER_REGISTRY_HTTP_CONNECTOR_B} &
  startNexus c ${PORT_NEXUS_C} ${DOCKER_REGISTRY_HTTP_CONNECTOR_C} &
  wait

  PASSWORD_NEXUS_A=$(cat ${DOWNLOAD_LOCATION_PASS}/nexus-a.txt)
  readonly PASSWORD_NEXUS_A
  PASSWORD_NEXUS_B=$(cat ${DOWNLOAD_LOCATION_PASS}/nexus-b.txt)
  readonly PASSWORD_NEXUS_B
  PASSWORD_NEXUS_C=$(cat ${DOWNLOAD_LOCATION_PASS}/nexus-c.txt)
  readonly PASSWORD_NEXUS_C
}

artifact() {
  mkdir -p "maven-releases/some/group${1}/File_${1}/1.0.0-2"
  echo someContent >"maven-releases/some/group${1}/File_${1}/1.0.0-2/File_${1}-1.0.0-2.jar"
  echo someContentZIP >"maven-releases/some/group${1}/File_${1}/1.0.0-2/File_${1}-1.0.0-2.zip"
  echo -e "<project>\n<modelVersion>4.0.0</modelVersion>\n<groupId>some.group${1}</groupId>\n<artifactId>File_${1}</artifactId>\n<version>1.0.0-2</version>\n</project>" >"maven-releases/some/group${1}/File_${1}/1.0.0-2/File_${1}-1.0.0-2.pom"
}

files() {
  for a in $(seq 100); do artifact "${a}"; done
}

upload() {
  echo " #134 archetype-catalog download issue"
  echo "URL: '${URL_NEXUS_A}/repository/maven-releases/archetype-catalog.xml'"
  echo "does not seem to contain a Maven artifact"
  curl -f ${URL_NEXUS_A}/repository/maven-releases/archetype-catalog.xml

  echo "Testing upload..."
  ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS_A} repositoriesV2 \
    --upload \
    --n3drRepo maven-releases \
    --directory-prefix "${PWD}" \
    --https=false
  echo
}

createHostedAPT() {
  echo "Creating apt repo..."
  curl "${2}/service/rest/beta/repositories/apt/hosted" \
    -s \
    -f \
    -u "admin:${1}" \
    -H "accept: application/json" \
    -H "Content-Type: application/json" \
    -d "{\"name\":\"REPO_NAME_HOSTED_APT\",\"online\":true,\"proxy\":{\"remoteUrl\":\"http://nl.archive.ubuntu.com/ubuntu/\"},\"storage\":{\"blobStoreName\":\"default\",\"strictContentTypeValidation\":true,\"writePolicy\":\"ALLOW_ONCE\"},\"apt\": {\"distribution\": \"bionic\"},\"aptSigning\": {\"keypair\": \"${N3DR_APT_GPG_SECRET}\",\"passphrase\": \"abc\"}}"
}

createHostedNPM() {
  echo "Creating npm repo..."
  curl "${2}/service/rest/v1/repositories/npm/hosted" \
    -s \
    -f \
    -u "admin:${1}" \
    -H "accept: application/json" \
    -H "Content-Type: application/json" \
    -d "{\"name\":\"REPO_NAME_HOSTED_NPM\",\"online\":true,\"storage\":{\"blobStoreName\":\"default\",\"strictContentTypeValidation\":true,\"writePolicy\":\"ALLOW_ONCE\"}}"
}

createHostedYum() {
  echo "Creating hosted yum repository..."
  ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS} configRepository \
    --configRepoName ${HOSTED_REPO_YUM} \
    --configRepoType yum \
    -p "${1}" \
    -n "${2}" \
    --https=false
}

createHostedDocker() {
  echo "Creating hosted docker repository..."
  ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS} configRepository \
    --configRepoName ${HOSTED_REPO_DOCKER} \
    --configRepoType docker \
    -p "${1}" \
    -n "${2}" \
    --https=false \
    --configRepoDockerPort="${DOCKER_REGISTRY_HTTP_CONNECTOR_INTERNAL}"
}

uploadDocker() {
  if [ "${NEXUS_API_VERSION}" != "beta" ]; then
    local docker_registry_uri=repository/testi/utrecht/n3dr
    createHostedDocker "${PASSWORD_NEXUS_A}" "${URL_NEXUS_A_V2}"

    echo "Testing docker upload..."
    docker login ${DOCKER_URL}:${DOCKER_REGISTRY_HTTP_CONNECTOR_A} \
      -p "${PASSWORD_NEXUS_A}" -u admin

    for d in $(seq 5); do
      local docker_registry_tag=localhost:${DOCKER_REGISTRY_HTTP_CONNECTOR_A}/${docker_registry_uri}:6.${d}.0
      docker pull "utrecht/n3dr:6.${d}.0"
      docker tag "utrecht/n3dr:6.${d}.0" "${docker_registry_tag}"
      docker push "${docker_registry_tag}"
    done

    echo
  else
    echo "docker upload not supported in beta API"
  fi
}

uploadYumArtifact() {
  curl -X 'POST' \
    ${URL_NEXUS_A}/service/rest/v1/components?repository=${HOSTED_REPO_YUM} \
    -s \
    -f \
    -u "admin:${PASSWORD_NEXUS_A}" \
    -H 'accept: application/json' \
    -H 'Content-Type: multipart/form-data' \
    -F "yum.asset=@${1};type=application/x-rpm" \
    -F "yum.asset.filename=${1}"
}

uploadYum() {
  if [ "${NEXUS_API_VERSION}" != "beta" ]; then
    createHostedYum "${PASSWORD_NEXUS_A}" "${URL_NEXUS_A_V2}"

    mkdir ${HOSTED_REPO_YUM}
    cd ${HOSTED_REPO_YUM}
    for i in $(seq 5 9); do
      curl -sL "https://yum.puppet.com/puppet-release-el-${i}.noarch.rpm" \
        -o "puppet${i}.rpm"
      uploadYumArtifact "puppet${i}.rpm"
    done

    cd ..
    echo
  else
    echo "yum upload not supported in beta API"
  fi
}

backupHelper() {
  count_downloads 301 "${1}"
  test_zip 132 "${1}"

  cleanup_downloads
}

backupRegexHelper() {
  count_downloads 5 "${1}"
  test_zip 4 "${1}"

  cleanup_downloads
}

anonymous() {
  echo "Testing backup by anonymous user..."
  local downloadDir="${DOWNLOAD_LOCATION}/anonymous/"
  ${N3DR_DELIVERABLE} repositoriesV2 \
    --backup \
    -n "${URL_NEXUS_A_V2}" \
    --n3drRepo maven-releases \
    -v "${NEXUS_API_VERSION}" \
    -z \
    --anonymous \
    --directory-prefix="${downloadDir}" \
    --directory-prefix-zip="${downloadDir}" \
    --https=false
  backupHelper "${downloadDir}"
}


backup() {
  echo "Testing backup..."
  local downloadDir="${DOWNLOAD_LOCATION}/backup/"
  ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS_A} repositoriesV2 \
    --backup \
    --n3drRepo maven-releases \
    -z \
    --directory-prefix="${downloadDir}" \
    --directory-prefix-zip="${downloadDir}" \
    --https=false
  backupHelper "${downloadDir}"
}

clean() {
  for r in a b c; do docker stop nexus-${r} || true; done
  docker stop rproxy-nginx-nexus3 || true
  cleanup_downloads
}

count_downloads() {
  local actual
  actual=$(find "${2}" -type f | wc -l)
  echo "Expected number of artifacts: ${1}"
  echo "Actual number of artifacts: ${actual}"
  echo "${actual}" | grep "${1}"
}

test_zip() {
  local size
  size=$(du "${2}"*n3dr-backup-*zip)
  echo "Actual ZIP size: ${size}"
  echo "Expected ZIP size: ${1}"
  echo "${size}" | grep "^${1}"
}

cleanup_downloads() {
  rm -rf nuget-hosted
  rm -rf REPO_NAME_HOSTED_APT
  rm -rf REPO_NAME_HOSTED_NPM
  rm -rf ${HOSTED_REPO_YUM}
  rm -rf maven-releases
  rm -rf "${DOWNLOAD_LOCATION}"
  rm -rf "${DOWNLOAD_LOCATION_PASS}"
  rm -rf "${DOWNLOAD_LOCATION_RPROXY}"
  rm -rf "${DOWNLOAD_LOCATION_SYNC}"
  rm -rf "${DOWNLOAD_LOCATION_SYNC_A}"
  rm -rf "${DOWNLOAD_LOCATION_SYNC_B}"
  rm -rf "${DOWNLOAD_LOCATION_SYNC_C}"
  rm -f n3dr-backup-*zip
  rm -f helloZip*zip
}

version() {
  echo "Check whether ./n3dr (N3DR_DELIVERABLE: ${N3DR_DELIVERABLE}) --version returns version"
  ${N3DR_DELIVERABLE} --version
  echo
}

rproxy() {
  if [ "${NEXUS_API_VERSION}" != "beta" ]; then
    local rproxy_conf=../../test/rproxy-nginx-nexus3.conf
    local rproxy_conf_tmp="${rproxy_conf}.tmp"

    echo "Testing rproxy in front of a nexus server..."
    ip_nexus_a=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' nexus-a)
    mkdir -p ${DOWNLOAD_LOCATION_RPROXY}
    sed -e "s|WILL_BE_REPLACED|${ip_nexus_a}|" "${rproxy_conf}" >"${rproxy_conf_tmp}"
    docker run -d --rm --name rproxy-nginx-nexus3 -p 9990:80 -v "${PWD}"/"${rproxy_conf_tmp}":/etc/nginx/nginx.conf nginx:1.21.5-alpine
    sleep 10
    ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS} repositoriesV2 \
      --count \
      --directory-prefix "${DOWNLOAD_LOCATION_RPROXY}" \
      -p "${PASSWORD_NEXUS_A}" \
      -n localhost:9990 \
      --https=false \
      --basePathPrefix=alternativeBasePathNexus3
  else
    echo "Rproxy check skipped in conjunction with beta API"
  fi
}

count() {
  echo "Counting artifacts..."
  ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS} count \
    -p "${PASSWORD_NEXUS_A}" \
    -n "${URL_NEXUS_A_V2}" \
    --https=false | grep 1500
}

countCSV() {
  local f=/tmp/helloworld

  echo "Counting artifacts and write to a CSV file..."
  ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS} count \
    -p "${PASSWORD_NEXUS_A}" \
    -n "${URL_NEXUS_A_V2}" \
    --https=false \
    --csv "${f}"
  du "${f}.csv" | grep 1488
}

main() {
  validate
  build
  nexus
  files

  export PASSWORD=${PASSWORD_NEXUS_A}
  readonly N3DR_DELIVERABLE_WITH_BASE_OPTIONS="${N3DR_DELIVERABLE} -u admin --showLogo=false"
  readonly N3DR_DELIVERABLE_WITH_BASE_OPTIONS_A="${N3DR_DELIVERABLE_WITH_BASE_OPTIONS} -p ${PASSWORD} -n ${URL_NEXUS_A_V2} -v ${NEXUS_API_VERSION}"
  upload
  anonymous
  backup
  # uploadDeb
  uploadDocker
  # uploadNPM
  # uploadNuget
  uploadYum
  # repositories
  # regex
  # zipName
  version
  rproxy
  count
  countCSV
  bats --tap ../../test/tests.bats
  echo "
In order to debug, issue:
N3DR_CLEAN_IN_CASE_OF_SUCCESS_OR_FAILURE=false ./test/integration-tests.sh and
login to ${URL_NEXUS_A}, ${URL_NEXUS_B} or ${URL_NEXUS_C} and login as admin
and respectively ${PASSWORD_NEXUS_A}, ${PASSWORD_NEXUS_B} or
${PASSWORD_NEXUS_C}"
}

if "${N3DR_CLEAN_IN_CASE_OF_SUCCESS_OR_FAILURE}"; then
  trap clean EXIT
fi

main
