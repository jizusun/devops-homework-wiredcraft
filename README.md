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

## 

## References
- https://www.docker.com/blog/multi-arch-build-and-images-the-simple-way/"I wonder", he said to himself, "what's in a book while it's closed.  Oh, I
know it's full of letters printed on paper, but all the same, something must
be happening, because as soon as I open it, there's a whole story with people
I don't know yet and all kinds of adventures and battles."
		-- Bastian B. Bux
You will visit the Dung Pits of Glive soon.
