Param()

# Determine script and project directories
$ScriptDir   = Split-Path -Parent $MyInvocation.MyCommand.Path
$ProtoDir    = Join-Path $ScriptDir '..\pb'
$OutDir      = Join-Path $ScriptDir '..\model'

# Derive your Go module path from go.mod
$ModulePath  = (go list -m).Trim()

# Create output directory if it doesn't exist
if (-not (Test-Path $OutDir)) {
    New-Item -ItemType Directory -Path $OutDir | Out-Null
}

# Build import-mapping flags for all your protos
$MFLAGS = @(
    "Mbinance.proto=${ModulePath}/model;model",
    "Mcoinbase.proto=${ModulePath}/model;model",
    "Mkraken.proto=${ModulePath}/model;model",
    "Mtrade.proto=${ModulePath}/model;model"
) -join ','

# Navigate to the proto directory
Push-Location $ProtoDir
Write-Host "Generating protobufs from '$ProtoDir' into '$OutDir' (module: $ModulePath)"

# Compile each .proto into the flat model/ directory with override mappings
Get-ChildItem -Filter '*.proto' | ForEach-Object {
    $file = $_.Name
    Write-Host "🛠 protoc → $file"
    & protoc `
        -I $ProtoDir `
        "--go_out=$MFLAGS,paths=source_relative:$OutDir" `
        "$file"

    if ($LASTEXITCODE -ne 0) {
        Write-Error "protoc failed on $file"
        Pop-Location
        exit $LASTEXITCODE
    }
}

# Return to original directory
Pop-Location
Write-Host "✅ Protobufs generated successfully into '$OutDir'."
