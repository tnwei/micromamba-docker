# micromamba-docker

Dockerfile for using micromamba for Python apps, which is the micro version of [`mamba`](https://github.com/mamba-org/mamba), a drop-in `conda` replacement. To run, clone this repo, drop your `conda` environment file here (output of `conda env export --from-history`), and specify file name at build time, e.g.:
```
docker build . --build-arg ENVFILE=example-environment.yml -t micromamba-docker
```

## Motivation

Micromamba is a suitable `conda` replacement for containers as no installation is needed, just need to download the binary. It is however still early in development, and does not have the feature to rig up `conda` envs directly from `environment.yml` yet. Thus this repo to set up a boilerplate Dockerfile, as well as a simple Go binary to convert the env file into an equivalent CLI command.

## Reference

Used this [gist](https://gist.github.com/wolfv/fe1ea521979973ab1d016d95a589dcde) as reference for micromamba commands.
