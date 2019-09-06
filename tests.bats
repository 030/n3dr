@test "invoking n3dr without password specification prints an error" {
  run go run main.go repositories -n http://localhost:9999 -u admin -b
  [ "$status" -eq 1 ]
  echo $output
  regex='.*Empty password. Verify whether the .n3drPass..*has been defined in ~/.n3dr.yaml'
  [[ "$output" =~ $regex ]]
}

@test "invoking n3dr with incorrect password specification prints an error" {
  run go run main.go repositories -n http://localhost:9999 -u admin -p INCORRECT_PASSWORD -b
  [ "$status" -eq 1 ]
  echo $output
  regex='.*ResponseCode: .401. and Message .401 Unauthorized. for URL: http://localhost:9999/service/rest/v1/repositories'
  [[ "$output" =~ $regex ]]
}