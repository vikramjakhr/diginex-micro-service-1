# Diginext Micro Service 1

## QuickStart

```bash
$ helm install ./diginex-micro-service-1 --name app1 --namespace dev
```

## Introduction

This chart bootstraps diginex-micro-service-1 deployment and service on a Kubernetes cluster using the Helm Package manager.

## Prerequisites

- Kubernetes 1.4+

## Installing the Chart

To install the chart with the release name `app1`:

```bash
$ helm install --name app1 ./diginex-micro-service-1
```

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `app1` deployment:

```bash
$ helm delete app1 --purge
```

The command removes all the Kubernetes components associated with the chart and deletes the release.