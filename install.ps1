$ErrorActionPreference = "Stop"

$repoUrl = "https://github.com/Jhonatan-Code-dev/go-promt-skills-rules.git"
$tempDir = Join-Path $env:TEMP ([System.Guid]::NewGuid().ToString())

try {
    Write-Host "[INFO] Descargando e instalando Go Rules & Skills..." -ForegroundColor Cyan

    git clone --depth 1 $repoUrl $tempDir 2>$null

    if (-not (Test-Path ".agents")) {
        New-Item -ItemType Directory -Path ".agents" | Out-Null
    }

    Copy-Item -Path "$tempDir\.agents\*" -Destination ".agents" -Recurse -Force

    Write-Host "[OK] Instalacion completada exitosamente." -ForegroundColor Green
    Write-Host "[INFO] Las reglas y skills se han instalado en .agents\" -ForegroundColor Yellow
}
finally {
    if (Test-Path $tempDir) {
        Remove-Item -Path $tempDir -Recurse -Force -ErrorAction SilentlyContinue
    }
}
