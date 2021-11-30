#!/bin/bash -e

if [ -z "${N3DR_APT_GPG_SECRET}" ]; then
  echo "N3DR_APT_GPG_SECRET should not be empty"
  echo "Create one by running:"
  echo "docker run -v /tmp/gpg-output:/root/.gnupg -v ${PWD}/test/gpg/:/tmp/ --rm -it vladgh/gpg --batch --generate-key /tmp/generate"
  echo "docker run --rm -it -v /tmp/gpg-output:/root/.gnupg -v ${PWD}/test/gpg/:/tmp/ vladgh/gpg --output /tmp/my_rsa_key --armor --export-secret-key joe@foo.bar"
  echo "Enter 'abc' as a password, if the prompt appears"
  echo "export N3DR_APT_GPG_SECRET=\$(sudo cat test/gpg/my_rsa_key | docker run -i m2s:2020-08-05)"
  echo "sudo rm -r /tmp/gpg-output"
  echo "rm test/gpg/my_rsa_key"
  echo
  printf "Note: Spaces and enters have to be escaped, i.e. '\\\n'->'\\\\\\\n' and ' '->'\ ' if the token is used in travis."
  exit 1
fi

if [ -z "${NEXUS_VERSION}" ]; then
  echo "NEXUS_VERSION empty, setting it to the default value"
  NEXUS_VERSION=3.37.0
fi

if [ -z "${NEXUS_API_VERSION}" ]; then
  echo "NEXUS_API_VERSION empty, setting it to the default value"
  NEXUS_API_VERSION=v1
fi

if [ -z "${N3DR_DELIVERABLE}" ]; then
  echo "N3DR_DELIVERABLE empty, setting it to the default value"
  N3DR_DELIVERABLE=n3dr
fi

readonly DOWNLOAD_LOCATION=/tmp/n3dr
readonly DOWNLOAD_LOCATION_PASS=${DOWNLOAD_LOCATION}-pass
readonly DOWNLOAD_LOCATION_SYNC=${DOWNLOAD_LOCATION}-sync
readonly PORT_NEXUS_A=9999
readonly PORT_NEXUS_B=9998
readonly PORT_NEXUS_C=9997
readonly URL_NEXUS_A=http://localhost:${PORT_NEXUS_A}
readonly URL_NEXUS_B=http://localhost:${PORT_NEXUS_B}
readonly URL_NEXUS_C=http://localhost:${PORT_NEXUS_C}
readonly URL_NEXUS_A_V2="${URL_NEXUS_A/http:\/\//}"
readonly URL_NEXUS_B_V2="${URL_NEXUS_B/http:\/\//}"
readonly URL_NEXUS_C_V2="${URL_NEXUS_C/http:\/\//}"

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
  source ./scripts/build.sh
  cd cmd/n3dr
}

startNexus() {
  # shellcheck disable=SC1091
  # as nexus-docker.sh is retrieved remotely
  source ./start.sh "${NEXUS_VERSION}" "${NEXUS_API_VERSION}" "nexus-${1}" "${2}" "${DOWNLOAD_LOCATION_PASS}" &>/dev/null
}

nexus() {
  curl -sL https://gist.githubusercontent.com/030/666c99d8fc86e9f1cc0ad216e0190574/raw/d6e02230d86bce0ad555fdb0330ed4162c735d09/nexus-docker.sh -o start.sh
  chmod +x start.sh

  startNexus a ${PORT_NEXUS_A} &
  startNexus b ${PORT_NEXUS_B} &
  startNexus c ${PORT_NEXUS_C} &
  wait

  PASSWORD_NEXUS_A=$(cat ${DOWNLOAD_LOCATION_PASS}/nexus-a.txt)
  readonly PASSWORD_NEXUS_A
  PASSWORD_NEXUS_B=$(cat ${DOWNLOAD_LOCATION_PASS}/nexus-b.txt)
  readonly PASSWORD_NEXUS_B
  PASSWORD_NEXUS_C=$(cat ${DOWNLOAD_LOCATION_PASS}/nexus-c.txt)
  readonly PASSWORD_NEXUS_C
}

artifact() {
  mkdir -p "maven-releases/some/group${1}/file${1}/1.0.0-2"
  echo someContent >"maven-releases/some/group${1}/file${1}/1.0.0-2/f.i-l.e.${1}-1.0.0-2.jar"
  echo someContentZIP >"maven-releases/some/group${1}/file${1}/1.0.0-2/file${1}-1.0.0-2.zip"
  echo -e "<project>\n<modelVersion>4.0.0</modelVersion>\n<groupId>some.group${1}</groupId>\n<artifactId>file${1}</artifactId>\n<version>1.0.0-2</version>\n</project>" >"maven-releases/some/group${1}/file${1}/1.0.0-2/file${1}-1.0.0-2.pom"
}

files() {
  for a in $(seq 100); do artifact "${a}"; done
}

upload() {
  echo "#134 archetype-catalog download issue"
  echo "URL:"
  echo "${URL_NEXUS_A}/repository/maven-releases/archetype-catalog.xml"
  echo "does not seem to contain a Maven artifact"
  curl -f ${URL_NEXUS_A}/repository/maven-releases/archetype-catalog.xml

  echo "Testing upload..."
  ./"${N3DR_DELIVERABLE}" upload -u admin \
    -p "${PASSWORD}" \
    -r maven-releases \
    -n ${URL_NEXUS_A} \
    -v "${NEXUS_API_VERSION}"
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

uploadDeb() {
  if [ "${NEXUS_API_VERSION}" != "beta" ]; then
    createHostedAPT "${PASSWORD_NEXUS_A}" "${URL_NEXUS_A}"

    mkdir REPO_NAME_HOSTED_APT
    cd REPO_NAME_HOSTED_APT
    curl -sL https://github.com/030/a2deb/releases/download/1.0.0/a2deb_1.0.0-0.deb -o a2deb.deb
    curl -sL https://github.com/030/n3dr/releases/download/5.0.1/n3dr_5.0.1-0.deb -o n3dr.deb
    curl -sL https://github.com/030/informado/releases/download/1.4.0/informado_1.4.0-0.deb -o informado.deb
    cd ..

    echo "Testing deb upload..."
    ./"${N3DR_DELIVERABLE}" upload -u=admin -p="${PASSWORD}" \
      -r=REPO_NAME_HOSTED_APT \
      -n=${URL_NEXUS_A} -v="${NEXUS_API_VERSION}" \
      -t=apt
    echo
  else
    echo "Deb upload not supported in beta API"
  fi
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

uploadNPM() {
  if [ "${NEXUS_API_VERSION}" != "beta" ]; then
    createHostedNPM "${PASSWORD_NEXUS_A}" "${URL_NEXUS_A}"

    mkdir REPO_NAME_HOSTED_NPM
    cd REPO_NAME_HOSTED_NPM
    curl https://registry.npmjs.org/@babel/core/-/core-7.12.10.tgz -o babel-core.tgz
    cd ..

    echo "Testing NPM upload..."
    ./"${N3DR_DELIVERABLE}" upload -u=admin -p="${PASSWORD}" -r=REPO_NAME_HOSTED_NPM \
      -n=${URL_NEXUS_A} -v="${NEXUS_API_VERSION}" \
      -t=npm
    echo
  else
    echo "NPM upload not supported in beta API"
  fi
}

uploadNuget() {
  if [ "${NEXUS_API_VERSION}" != "beta" ]; then
    mkdir nuget-hosted
    cd nuget-hosted
    curl -sL https://chocolatey.org/api/v2/package/n3dr/5.2.6 -o n3dr.5.2.6.nupkg
    cd ..

    echo "Testing nuget upload..."
    ./"${N3DR_DELIVERABLE}" upload -u=admin -p="${PASSWORD}" -r=nuget-hosted \
      -n=${URL_NEXUS_A} -v="${NEXUS_API_VERSION}" \
      -t=nuget
    echo
  else
    echo "Nuget upload not supported in beta API"
  fi
}

backupHelper() {
  if [ "${NEXUS_VERSION}" == "3.9.0" ]; then
    count_downloads 300
    test_zip 152
  else
    count_downloads 400
    test_zip 196
  fi
  cleanup_downloads
}

anonymous() {
  echo "Testing backup by anonymous user..."
  ./"${N3DR_DELIVERABLE}" backup -n ${URL_NEXUS_A} -r maven-releases -v "${NEXUS_API_VERSION}" -z --anonymous
  backupHelper
}

backup() {
  echo "Testing backup..."
  ./"${N3DR_DELIVERABLE}" backup -n ${URL_NEXUS_A} -u admin -p "${PASSWORD}" -r maven-releases -v "${NEXUS_API_VERSION}" -z
  backupHelper
}

regex() {
  echo "Testing backup regex..."
  ./"${N3DR_DELIVERABLE}" backup -n ${URL_NEXUS_A} -u admin -p "${PASSWORD}" -r maven-releases -v "${NEXUS_API_VERSION}" -x 'some/group42' -z
  if [ "${NEXUS_VERSION}" == "3.9.0" ]; then
    count_downloads 3
    test_zip 4
  else
    count_downloads 4
    test_zip 4
  fi
  cleanup_downloads
  echo -e "\nTesting repositories regex..."
  ./"${N3DR_DELIVERABLE}" repositories -n ${URL_NEXUS_A} -u admin -p "${PASSWORD}" -v "${NEXUS_API_VERSION}" -b -x 'some/group42' -z
  if [ "${NEXUS_VERSION}" == "3.9.0" ]; then
    count_downloads 3
    test_zip 4
  else
    count_downloads 4
    test_zip 4
  fi
  cleanup_downloads
}

repositories() {
  local cmd="./$N3DR_DELIVERABLE repositories -n ${URL_NEXUS_A} -u admin -p ${PASSWORD} -v ${NEXUS_API_VERSION}"

  echo "Testing repositories..."
  $cmd -a | grep maven-releases

  echo "> Counting number of repositories..."
  expected_number=7
  if [ "${NEXUS_API_VERSION}" == "beta" ]; then
    expected_number=5
  fi
  actual_number="$($cmd -c | tail -n1)"
  echo -n "Number of repositories. Expected: ${expected_number}. Actual: ${actual_number}"
  [ "${actual_number}" == "${expected_number}" ]

  echo "> Testing zip functionality..."
  testZipSizeDir=/tmp/n3dr/testZipSize/
  $cmd -b -z \
    --directory-prefix ${testZipSizeDir} \
    --directory-prefix-zip ${testZipSizeDir}
  if [ "${NEXUS_VERSION}" == "3.9.0" ]; then
    count_downloads 301
    test_zip 132 ${testZipSizeDir}
  else
    count_downloads 402
    test_zip 208 ${testZipSizeDir}
  fi
  cleanup_downloads
}

zipName() {
  echo "Testing zipName..."
  ./"${N3DR_DELIVERABLE}" backup -n=${URL_NEXUS_A} -u=admin -p="${PASSWORD}" -r=maven-releases -v="${NEXUS_API_VERSION}" -z -i=helloZipFile.zip
  ./"${N3DR_DELIVERABLE}" repositories -n ${URL_NEXUS_A} -u admin -p "${PASSWORD}" -v "${NEXUS_API_VERSION}" -b -z -i=helloZipRepositoriesFile.zip
  find . -name "helloZip*" -type f | wc -l | grep 2
}

clean() {
  for r in a b; do NEXUS_DOCKER_NAME=nexus-${r} cleanup; done
  cleanup_downloads
}

count_downloads() {
  local actual
  actual=$(find ${DOWNLOAD_LOCATION} -type f | wc -l)
  echo "Expected number of artifacts: ${1}"
  echo "Actual number of artifacts: ${actual}"
  echo "${actual}" | grep "${1}"
}

test_zip() {
  local size
  size=$(du "${2}"n3dr-backup-*zip)
  echo "Actual ZIP size: ${size}"
  echo "Expected ZIP size: ${1}"
  echo "${size}" | grep "^${1}"
}

cleanup_downloads() {
  rm -rf nuget-hosted
  rm -rf REPO_NAME_HOSTED_APT
  rm -rf REPO_NAME_HOSTED_NPM
  rm -rf maven-releases
  rm -rf "${DOWNLOAD_LOCATION}"
  rm -rf "${DOWNLOAD_LOCATION_SYNC}"
  rm -f n3dr-backup-*zip
  rm -f helloZip*zip
}

version() {
  echo "Check whether ./n3dr (N3DR_DELIVERABLE: ${N3DR_DELIVERABLE}) --version returns version"
  "./${N3DR_DELIVERABLE}" --version
  echo
}

cac() {
  echo "Configuration as code"
}

sync() {
  if [ "${NEXUS_API_VERSION}" != "beta" ]; then
    echo "Testing the sync between NexusA and followers NexusB and C..."
    createHostedAPT "${PASSWORD_NEXUS_B}" "${URL_NEXUS_B}"
    createHostedAPT "${PASSWORD_NEXUS_C}" "${URL_NEXUS_C}"
    createHostedNPM "${PASSWORD_NEXUS_B}" "${URL_NEXUS_B}"
    createHostedNPM "${PASSWORD_NEXUS_C}" "${URL_NEXUS_C}"

    ./"${N3DR_DELIVERABLE}" sync \
      --otherNexus3Passwords="${PASSWORD_NEXUS_B}","${PASSWORD_NEXUS_C}" \
      --otherNexus3Users=admin,admin \
      --otherNexus3URLs="${URL_NEXUS_B_V2}","${URL_NEXUS_C_V2}" \
      --directory-prefix "${DOWNLOAD_LOCATION_SYNC}" \
      -p "${PASSWORD_NEXUS_A}" \
      -n "${URL_NEXUS_A_V2}" \
      -u admin
  else
    echo "RepositoriesV2 sync not supported in beta API"
  fi
}

main() {
  validate
  build
  nexus
  files

  export PASSWORD=${PASSWORD_NEXUS_A}
  upload
  anonymous
  backup
  uploadDeb
  uploadNPM
  uploadNuget
  repositories
  regex
  zipName
  version
  sync
  bats --tap ../../test/tests.bats
  echo "
In order to debug, comment out the 'trap clean EXIT', run this script again and
login to ${URL_NEXUS_A}, ${URL_NEXUS_B} or ${URL_NEXUS_C} and login as admin
and respectively ${PASSWORD_NEXUS_A}, ${PASSWORD_NEXUS_B} or
${PASSWORD_NEXUS_C}"
}

# trap clean EXIT
main
