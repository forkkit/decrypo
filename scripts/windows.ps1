param(
    [Parameter(Mandatory=$True, Position=0, ValueFromPipeline=$false)]
    [System.String]
    $TravisOS,

    [Parameter(Mandatory=$True, Position=1, ValueFromPipeline=$false)]
    [System.String]
    $TravisArch,
    
	[Parameter(Mandatory=$True, Position=2, ValueFromPipeline=$false)]
    [System.String]
    $TravisTag
)

$Time = Get-Date -Format "yyyy-MM-ddTHH:mm:ssK"
$Name = "$($TravisOS)_$($TravisArch)"

New-Item "dist/$($Name)" -ItemType Directory

$LDFlags = "-s -w -X github.com/ajdnik/decrypo/build.version=$($TravisTag) -X github.com/ajdnik/decrypo/build.datetime=$($Time)"
$Output = "dist/$($Name)/decrypo.exe"

Set-Item -Path Env:CGO_ENABLED -Value 1

& go build -ldflags $LDFlags -o $Output

Add-Type -assembly "system.io.compression.filesystem"
[io.compression.zipfile]::CreateFromDirectory("dist/$($Name)", "dist/$($Name).zip")

Get-FileHash "dist/$($Name).zip" -Algorithm SHA256 | Export-Clixml -Path "dist/$($Name).sha256"

Remove-Item -Path "dist/$($Name)" -Recurse
