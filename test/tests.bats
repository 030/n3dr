@test "invoking n3dr with incorrect password specification prints an error" {
  run ./${N3DR_DELIVERABLE} repositories -n http://localhost:9999 -u admin -p INCORRECT_PASSWORD -b
  [ "$status" -eq 1 ]
  echo $output
  regex='.*ResponseCode: .401. and Message .401 Unauthorized. for URL: http://localhost:9999/service/rest/v1/repositories'
  [[ "$output" =~ $regex ]]
}

@test "invoking n3dr with unreachable URL exits after 6 attempts" {
  run ./${N3DR_DELIVERABLE} repositories -n http://localhost:99999 -u admin -p INCORRECT_PASSWORD -b
  [ "$status" -eq 1 ]
  echo $output
  regex='.*http://localhost:99999/service/rest/v1/repositories giving up after 6 attempt\(s\)'
  [[ "$output" =~ $regex ]]
}

@test "invoking n3dr with repositories subcommand and incorrect URL exits" {
  run ./${N3DR_DELIVERABLE} repositories -n http://localhost:99999/ -u admin -p INCORRECT_PASSWORD -b
  [ "$status" -eq 1 ]
  echo $output
  regex=".*The Nexus3 URL seems to be incorrect. Ensure that it does not end with a '/'. Error: 'URL: regular expression mismatch'"
  [[ "$output" =~ $regex ]]
}

@test "invoking n3dr with backup subcommand and incorrect URL exits" {
  run ./${N3DR_DELIVERABLE} backup -u bla -n http://hihi/ -r bla -z
  [ "$status" -eq 1 ]
  echo $output
  regex=".*The Nexus3 URL seems to be incorrect. Ensure that it does not end with a '/'. Error: 'URL: regular expression mismatch'"
  [[ "$output" =~ $regex ]]
}

@test "invoking n3dr with upload subcommand and incorrect URL exits" {
  run ./${N3DR_DELIVERABLE} upload -u bla -n http://hihi/ -r bla
  [ "$status" -eq 1 ]
  echo $output
  regex=".*The Nexus3 URL seems to be incorrect. Ensure that it does not end with a '/'. Error: 'URL: regular expression mismatch'"
  [[ "$output" =~ $regex ]]
}

@test "invoking n3dr with version subcommand should return version" {
  run ./${N3DR_DELIVERABLE} --version
  [ "$status" -eq 0 ]
  echo $output
  regex=".*n3dr version.*"
  [[ "$output" =~ $regex ]]
}
