package main

import (
    "flag"
    "fmt"
    "io"
    "os"
    "path/filepath"
   
)

func main() {
    sourceDir := flag.String("source", "", "Source directory to backup")
    targetDir := flag.String("target", "", "Target directory where backup will be stored")
    flag.Parse()

    // Validate source and target directories
    if *sourceDir == "" || *targetDir == "" {
        fmt.Println("Source and target directories must be provided")
        return
    }

    // Start backup process
    err := backup(*sourceDir, *targetDir)
    if err != nil {
        fmt.Printf("Backup failed: %v\n", err)
        return
    }

    fmt.Println("Backup completed successfully")
}

func backup(sourceDir, targetDir string) error {
    return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Skip directories
        if info.IsDir() {
            return nil
        }

        relPath, err := filepath.Rel(sourceDir, path)
        if err != nil {
            return err
        }

        targetPath := filepath.Join(targetDir, relPath)
        targetInfo, err := os.Stat(targetPath)

        // If target file doesn't exist or is older, copy the file
        if os.IsNotExist(err) || targetInfo.ModTime().Before(info.ModTime()) {
            err := copyFile(path, targetPath)
            if err != nil {
                return err
            }
            fmt.Printf("Copied: %s\n", relPath)
        }

        return nil
    })
}

func copyFile(sourceFile, destFile string) error {
    source, err := os.Open(sourceFile)
    if err != nil {
        return err
    }
    defer source.Close()

    destDir := filepath.Dir(destFile)
    if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
        return err
    }

    destination, err := os.Create(destFile)
    if err != nil {
        return err
    }
    defer destination.Close()

    _, err = io.Copy(destination, source)
    return err
}
