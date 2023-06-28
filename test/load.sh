#!/bin/bash

# artifact() {
#   local dir="${FILES_LOCATION}/${2}/some/group${1}/File_${1}/1.0.0-2"
#   local file="${dir}/File_${1}-1.0.0-2"
#   mkdir -p "${dir}"
#   dd if=/dev/urandom of=${file} bs=1M count=${1}
#   cp ${file} "${file}.jar"
#   zip ${file}.zip ${file}
#   echo -e "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<project>\n<modelVersion>4.0.0</modelVersion>\n<groupId>some.group${1}</groupId>\n<artifactId>File_${1}</artifactId>\n<version>1.0.0-2</version>\n</project>" >"${file}.pom"
# }

# artifactWithoutPom() {
#   local dir="${FILES_LOCATION}/${2}/some/group${1}/File_${1}_without_pom/1.0.0-2"
#   local file="${dir}/File_${1}-1.0.0-2"
#   mkdir -p "${dir}"
#   dd if=/dev/urandom of=${file} bs=1M count=${1}
#   cp ${file} "${file}.jar"
#   zip ${file}.zip ${file}
# }

# createHostedMaven2() {
#   echo "Creating hosted Maven2 repository number: '${3}'..."
#   ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS} configRepository \
#     --configRepoName "maven-releases-${3}" \
#     --configRepoType maven2 \
#     -p "${1}" \
#     -n "${2}" \
#     --strictContentTypeValidation=false \
#     --https=false
# }

# load() {
#   for i in $(seq 25); do
#     files "maven-releases-${i}"
#     createHostedMaven2 "${PASSWORD_NEXUS_A}" "${URL_NEXUS_A_V2}" "${i}"
#   done

#   ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS_A} repositoriesV2 \
#     --upload \
#     --directory-prefix "${FILES_LOCATION}" \
#     --https=false
# }

# files() {
#   for a in $(seq 10); do artifact "${a}" "${1}"; done
# }

# files "maven-releases"
