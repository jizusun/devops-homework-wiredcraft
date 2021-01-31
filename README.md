# devops-test


## Getting Started 
```sh
git clone --recurse-submodules git@github.com:jizusun/devops-homework-wiredcraft.git
# checkout other branches into subfolders
git worktree add -B gh-pages public origin/gh-pages
git worktree add -B src src origin/src
```


## Set the hugo template

```sh
# https://gohugo.io/getting-started/quick-start/#step-2-create-a-new-site

cd ..
hugo new site devops-homework-wiredcraft --force
cd devops-homework-wiredcraft
git submodule add https://github.com/budparr/gohugo-theme-ananke.git themes/ananke

hugo new posts/my-first-post.md
mkdir -p layouts/partials/
cp themes/ananke/layouts/partials/site-footer.html layouts/partials/site-footer.html

/usr/games/fortune

```
