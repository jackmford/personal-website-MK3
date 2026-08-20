# personal-website-MK3

[jackmitchellfordyce.com](https://jackmitchellfordyce.com)

A Go website with embedded templates, static assets, and Markdown blog posts.

## Operate

### Run locally

```sh
make run
```

The site listens on `http://localhost:4000`. Embedded files are compiled into
the binary, so restart `make run` after changing templates, CSS, or posts.

```sh
go test ./...
```

### Publish a blog post

Add a Markdown file to `ui/content/blog/` using the front matter and structure
of an existing post. Posts are embedded and sorted by their `date` front
matter when the application starts.

### Release and deploy

1. Merge the website change to `main` after its checks pass.
2. Create and push a version tag from `main`, for example:

   ```sh
   git tag v3.0.6
   git push origin v3.0.6
   ```

   The tag-triggered GitHub Actions workflow tests the app and publishes a
   Linux `.deb` package.

3. In the separate `ansible` repository, set
   `roles/personal-website/defaults/main.yml` to that tag and run:

   ```sh
   ansible-playbook playbooks/personal_website.yml -u <ssh-user>
   ```

   The playbook installs the release package and restarts the
   `personal-website` systemd service behind Caddy.

### Verify production

```sh
curl -fsS https://jackmitchellfordyce.com/health
```

Check `https://jackmitchellfordyce.com/metrics` for application metrics.
Cloudflare Web Analytics provides visitor analytics after the deployed pages
receive traffic.
