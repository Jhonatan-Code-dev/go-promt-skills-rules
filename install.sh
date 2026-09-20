#!/usr/bin/env bash
set -e

REPO_URL="https://github.com/Jhonatan-Code-dev/go-promt-skills-rules.git"
TMP_DIR=$(mktemp -d)

cleanup() {
  rm -rf "$TMP_DIR"
}
trap cleanup EXIT

echo "[INFO] Descargando e instalando Go Rules & Skills..."

git clone --depth 1 "$REPO_URL" "$TMP_DIR" > /dev/null 2>&1

mkdir -p .agents

cp -r "$TMP_DIR/.agents/"* .agents/

echo "[OK] Instalacion completada exitosamente."
echo "[INFO] Las reglas y skills se han instalado en .agents/"
