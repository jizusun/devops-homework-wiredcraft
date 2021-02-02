# Wiredcraft-Hugo Pipeline

[![codecov](https://codecov.io/gh/jizusun/devops-homework-wiredcraft/branch/src/graph/badge.svg?token=N0R6ZOVKJ2)](https://codecov.io/gh/jizusun/devops-homework-wiredcraft)

## Getting Started with the development

```sh
git clone --recurse-submodules git@github.com:jizusun/devops-homework-wiredcraft.git
git checkout hugo
# https://gohugo.io/hosting-and-deployment/hosting-on-github/
# checkout branches into subfolders
git worktree add -B gh-pages public origin/gh-pages
git worktree add -B src src origin/src
# switch to the go source code
cd src
```

## Hugo templates, Ansible, and terraform

Check the [hugo](https://github.com/jizusun/devops-homework-wiredcraft/tree/hugo) branch
