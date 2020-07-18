@test "invoking n3dr with incorrect password specification prints an error" {
  run go run main.go repositories -n http://localhost:9999 -u admin -p INCORRECT_PASSWORD -b
  [ "$status" -eq 1 ]
  echo $output
  regex='.*ResponseCode: .401. and Message .401 Unauthorized. for URL: http://localhost:9999/service/rest/v1/repositories'
  [[ "$output" =~ $regex ]]
}

@test "invoking n3dr with unreachable URL exists after 6 attempts" {
  run go run main.go repositories -n http://localhost:99999 -u admin -p INCORRECT_PASSWORD -b
  [ "$status" -eq 1 ]
  echo $output
  regex='.*http://localhost:99999/service/rest/v1/repositories giving up after 6 attempts'
  [[ "$output" =~ $regex ]]
}