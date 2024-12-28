# Better Youtube Playlists Htmx

Buildling a better youtube playlist using htmx and a go backend.

## Infrastructure

There are two Cloudformation templates: one for s3 and one for lambda. When creating the infrastructure for the first time (a manual process), there are a couple of steps that need to be followed:

1. Configuring aws credentials using `aws configure` and setting the correct `AWS Access Key ID`, `AWS Secret Access Key` and `region name`.
2. Deploying the s3 bucket using

```
make deploy-s3
```

3. Now whenever a code change is pushed, your code should be deployed to EC2. You can visit the site via https://better-youtube-playlists.com.

If for any reason you need to delete the infrastructure, this can be done via this two commands:

```
make delete-s3-stack
make delete-ec2-stack
```
