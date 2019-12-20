# mkghurl

Make github url from OS env variables.

The tool reads the following values:

- $GITHUB_USER
- $GITHUB_TOKEN

and make url from the following url-pattern:

```html
https://$GITHUB_USER:$GITHUB_TOKEN@github.com/$GITHUB_USER/repo-name.git
```

## Install

```shell script
$ go install github.com/devlights/mkghurl
```

## Run

```shell script
$ mkghurl repo-name
```