# devops-test 

![Update Hugo (dev)](https://github.com/jizusun/devops-homework-wiredcraft/workflows/Update%20Hugo%20(dev)/badge.svg)

- http://dev.wiredcraft.edtechstar.com/
- http://staging.wiredcraft.edtechstar.com/
- https://github.com/jizusun/devops-homework-wiredcraft/actions
- https://github.com/Wiredcraft/test-devops


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

Check the [src](https://github.com/jizusun/devops-homework-wiredcraft/tree/src) branch