# Better Youtube Playlists Htmx

Buildling a better youtube playlist using htmx and a go backend.

## Infrastructure

There are two Cloudformation templates: one for s3 and one for lambda. When creating the infrastructure for the first time (a manual process), there are a couple of steps that need to be followed:

1. Configuring aws credentials using `aws configure` and setting the correct `AWS Access Key ID`, `AWS Secret Access Key` and `region name`.
2. Deploying the s3 bucket using

```
aws cloudformation deploy \
           --template-file infrastructure/s3-template.json \
           --stack-name  BetterYoutubePlaylistsS3Stack \
           --capabilities CAPABILITY_NAMED_IAM
```

3.  Manually uploading a placeholder zip file called `better-youtube-playlists-lambda.zip` to it, this should contain code in the shape of a default golang lambda handler.
4.  Deploying the lambda function using

```
aws cloudformation deploy \
           --template-file infrastructure/better-youtube-playlists-template.json \
           --stack-name  BetterYoutubePlaylistsLambdaStack \
           --capabilities CAPABILITY_NAMED_IAM
```

Now you should have an s3 bucket with a placeholder zip file and lambda infrastructure that references the s3 bucket zip file. You should now be able upload any custom lambda code and replace the placeholder zip file in the bucket which should now be executable by the lambda.

If for any reason you need to delete the infrastructure, this can be done via these two commands:

```
aws cloudformation delete-stack --stack-name BetterYoutubePlaylistsLambdaStack
aws cloudformation delete-stack --stack-name BetterYoutubePlaylistsS3Stack
```
