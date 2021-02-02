# devops-test 


## Set the hugo template

```sh
# https://gohugo.io/getting-started/quick-start/#step-2-create-a-new-site
hugo new site devops-homework-wiredcraft 
git submodule add https://github.com/budparr/gohugo-theme-ananke.git themes/ananke

hugo new posts/my-first-post.md
mkdir -p layouts/partials/
cp themes/ananke/layouts/partials/site-footer.html layouts/partials/site-footer.html
```


## References
- https://www.docker.com/blog/multi-arch-build-and-images-the-simple-way/
