# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Pulumi provider for generating Kubernetes configurations (kubeconfigs) for different cloud providers:
- **AKS (Azure Kubernetes Service)** - Component resource
- **EKS (Amazon Elastic Kubernetes Service)** - Custom resource  
- **GKE (Google Kubernetes Engine)** - Custom resource

The provider is built using the Pulumi Go Provider SDK and generates multi-language SDKs for Go, Node.js, Python, and .NET.

## Development Commands

All commands use Task (taskfile.yaml):

### Core Development
- `task build_provider` - Build the provider binary
- `task clean` - Remove build artifacts and schema
- `task ensure` - Run `go mod tidy`
- `task lint` - Run golangci-lint

### SDK Generation
- `task get_schema` - Generate schema.json from provider
- `task generate_sdks` - Generate all language SDKs (requires built provider)
- `task only_generate_sdks` - Generate SDKs without rebuilding provider
- `task generate_nodejs_sdk` - Generate only Node.js SDK
- `task generate_dotnet_sdk` - Generate only .NET SDK  
- `task generate_go_sdk` - Generate only Go SDK
- `task generate_python_sdk` - Generate only Python SDK

### Development Workflow
- `task watch` - Watch Go files and rebuild provider on changes

## Architecture

### Core Components
- `main.go` - Provider entry point and SDK configuration
- `pkg/kubeconfig.go` - Shared kubeconfig data structures
- `pkg/aks.go` - AKS component resource (uses Azure Native SDK)
- `pkg/eks.go` - EKS custom resource (generates AWS CLI-based config)
- `pkg/gke.go` - GKE custom resource (generates gcloud auth plugin config)

### Resource Types
- **AKS**: Component resource using Azure Native containerservice API
- **EKS/GKE**: Custom resources with Create/Diff/Update lifecycle methods

### Generated SDKs
All SDKs are generated in `sdk/` directory with language-specific subdirectories. Each includes proper package metadata and dependencies for the target ecosystem.

## Key Dependencies
- Pulumi Go Provider SDK
- Azure Native SDK (for AKS)
- golangci-lint (for linting)

The provider generates kubeconfigs with cloud-specific authentication methods:
- AKS: Uses Azure Native SDK to retrieve cluster credentials
- EKS: Generates config with AWS CLI `eks get-token` command
- GKE: Generates config with `gke-gcloud-auth-plugin` command