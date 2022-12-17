readonly NEXUS_URL_ERROR_MSG_REGEX=".*the Nexus3 URL seems to be incorrect. Verify that it complies to the regex that is defined in the 'Nexus3 Struct' and that it does not end with a '/'. Error: 'URL: regular expression mismatch'"

@test "invoking n3dr with version subcommand should return version" {
  run ./${N3DR_DELIVERABLE} --version
  [ "$status" -eq 0 ]
  echo $output
  regex=".*n3dr version.*"
  [[ "$output" =~ $regex ]]
}
