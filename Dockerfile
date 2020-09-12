# Dockerfile for utilizing micromamba
# Refer to docs at https://github.com/mamba-org/mamba

# Using Ubuntu as base image
FROM ubuntu:18.04

# Change shell from sh to bash 
SHELL ["/bin/bash", "-c"]

# Download wget, and download micromamba into /bin/
RUN apt-get -qq update && apt-get -qq -y install wget
RUN wget -qO- https://micromamba.snakepit.net/api/micromamba/linux-64/latest | tar -xvj bin/micromamba -C /bin/

# Sourcing env vars?
# Run below if using micromamba in Docker
ARG MAMBA_EXE=bin/micromamba
ARG MAMBA_ROOT_PREFIX=/micromamba-base
ENV MAMBA_EXE=$MAMBA_EXE
# ENV MAMBA_ROOT_PREFIX=$MAMBA_ROOT_PREFIX

# micromamba shell init
RUN micromamba shell init -s bash -p /micromamba-base

# Example: Making envs for project `alice`
# Note: Path of prefix needs to be absolute as per Jul 25, 2020 from
# https://github.com/TheSnakePit/mamba/issues/386
# Note: Channels need to be explicitly specified!
RUN source ~/.bashrc && micromamba create -y -p /alice/ pandas numpy seaborn -c defaults -y

# Example: Making envs for project `bob`
# Different channels require different lines for installation
RUN micromamba create -y -p /bob/ statsmodels scikit-learn tensorflow==2.2 requests beautifulsoup4 graphviz matplotlib pandas numpy seaborn -c conda-forge
