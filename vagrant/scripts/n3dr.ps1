function Build {
  Write-Output "Building N3DR..."
  Set-Location C:\vagrant\cmd\n3dr
  go build -buildvcs=false
}

function Main {
  Build
}

Main
