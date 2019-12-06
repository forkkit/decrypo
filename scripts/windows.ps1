$Time = Get-Date -Format "yyyy-MM-ddTHH:mm:ssK"
$TravisOS = Get-Variable TRAVIS_OS_NAME -valueOnly
$TravisArch = Get-Variable TRAVIS_CPU_ARCH -valueOnly
$TravisTag = Get-Variable TRAVIS_TAG -valueOnly

$Name = "$($TravisOS)_$($TravisArch)"

New-Item "dist/$($Name)" -ItemType Directory

$LDFlags = "-s -w -X github.com/ajdnik/decrypo/build.version=$($TravisTag) -X github.com/ajdnik/decrypo/build.datetime=$($Time)"
$Output = "dist/$($Name)/decrypo.exe"

& go build -ldflags $LDFlags -o $Output

Add-Type -assembly "system.io.compression.filesystem"
[io.compression.zipfile]::CreateFromDirectory("dist/$($Name)", "dist/$($Name).zip")

Get-FileHash "dist/$($Name).zip" -Algorithm SHA256 | Export-Clixml -Path "dist/$($Name).sha256"

Remove-Item -Path "dist/$($Name)" -Recurse
