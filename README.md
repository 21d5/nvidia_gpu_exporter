# nvidia_gpu_exporter

[![build](https://github.com/utkuozdemir/nvidia_gpu_exporter/actions/workflows/build.yml/badge.svg)](https://github.com/utkuozdemir/nvidia_gpu_exporter/actions/workflows/build.yml)
[![Coverage Status](https://coveralls.io/repos/github/utkuozdemir/nvidia_gpu_exporter/badge.svg?branch=master)](https://coveralls.io/github/utkuozdemir/nvidia_gpu_exporter?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/utkuozdemir/nvidia_gpu_exporter?kill_cache=1)](https://goreportcard.com/report/github.com/utkuozdemir/nvidia_gpu_exporter)
![Latest GitHub release](https://img.shields.io/github/release/utkuozdemir/nvidia_gpu_exporter.svg)
[![GitHub license](https://img.shields.io/github/license/utkuozdemir/nvidia_gpu_exporter)](https://github.com/utkuozdemir/nvidia_gpu_exporter/blob/master/LICENSE)
![GitHub all releases](https://img.shields.io/github/downloads/utkuozdemir/nvidia_gpu_exporter/total)
![Docker Pulls](https://img.shields.io/docker/pulls/utkuozdemir/nvidia_gpu_exporter)

Nvidia GPU exporter for prometheus, using `nvidia-smi` binary to gather metrics.

## Introduction

There are many Nvidia GPU exporters out there however they have problems such as not being maintained, 
not providing pre-built binaries, having a dependency to Linux and/or Docker, 
targeting enterprise setups (DCGM) and so on.

This is a simple exporter that uses `nvidia-smi(.exe)` binary to collect, parse and export metrics.
This makes it possible to run it on Windows and get GPU metrics while gaming - no Docker or Linux required.

This project is based on [a0s/nvidia-smi-exporter](https://github.com/a0s/nvidia-smi-exporter).
However, this one is written in Go to produce a single, static binary.

**If you are a gamer who's into monitoring, you are in for a treat.**

## Highlights

- Will work on any system that has `nvidia-smi(.exe)?` binary - Windows, Linux, MacOS... No C bindings required
- Doesn't even need to run the monitored machine: can be configured to execute `nvidia-smi` command remotely
- No need for a Docker or Kubernetes environment
- Auto-discovery of the metric fields `nvidia-smi` can expose (future-compatible)
- Comes with its own [Grafana dashboard](https://grafana.com/grafana/dashboards/14574)

## Visualization

You can use the official [Grafana dashboard](https://grafana.com/grafana/dashboards/14574)
to see your GPU metrics in a nicely visualized way.

Here's how it looks like:
![Grafana dashboard](https://raw.githubusercontent.com/utkuozdemir/nvidia_gpu_exporter/master/grafana/dashboard.png)


## Installation
See [INSTALL.md](INSTALL.md) for details.

## Configuration
See [CONFIGURE.md](CONFIGURE.md) for details.

## Metrics
See [METRICS.md](METRICS.md) for details.

## Contributing
See [CONTRIBUTING.md](CONTRIBUTING.md) for details.
