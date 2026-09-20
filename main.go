package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed .agents/*
var embeddedAgents embed.FS

func main() {
	fmt.Println("[INFO] Instalando Go Rules & Skills (.agents) en tu proyecto...")

	targetDir := "."
	agentsDir := filepath.Join(targetDir, ".agents")

	// Crear el directorio .agents si no existe
	if err := os.MkdirAll(agentsDir, 0755); err != nil {
		fmt.Printf("[ERROR] Error al crear directorio .agents: %v\n", err)
		os.Exit(1)
	}

	installedCount := 0

	err := fs.WalkDir(embeddedAgents, ".agents", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Determinar la ruta relativa de destino
		relPath, err := filepath.Rel(".agents", path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(agentsDir, relPath)

		if d.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		// Leer el archivo embebido
		data, err := embeddedAgents.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error al leer archivo embebido %s: %w", path, err)
		}

		// Escribir en el destino
		if err := os.WriteFile(destPath, data, 0644); err != nil {
			return fmt.Errorf("error al escribir archivo %s: %w", destPath, err)
		}

		fmt.Printf("  [OK] Copiado: %s\n", destPath)
		installedCount++
		return nil
	})

	if err != nil {
		fmt.Printf("[ERROR] Error durante la instalación: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n[EXITO] Instalacion completada exitosamente.")
	fmt.Printf("[INFO] Se instalaron %d archivos de reglas y skills en %s\n", installedCount, agentsDir)
	fmt.Println("       Tus agentes de IA (Antigravity, Cursor, Windsurf, etc.) ahora utilizaran automaticamente tus reglas de Go.")
}
