# tsubst
Substitute environment variables on GO templates

## Example - Basic Usage

Given the following "template.yml" file:
```yaml
job1:
  environment: {{.ENV}}
```
Run this:
```console
$ export ENV=prod
$ tsubst template.yml > prod.yml
$ cat prod.yml
job1:
  environment: prod
```

## Example - Conditionals

Given the following "template.yml" file:
```yaml
server-deployment:
  {{if eq .MODE "basic"}}
  instances: 1
  {{else}}
  instances: 3
  {{end}}
```
Run this:
```console
$ export MODE=basic
$ tsubst template.yml > basic.yml
$ cat basic.yml
server-deployment:
  instances: 1
```

## More information
For a complete reference on the template syntax and capabilities,
see the [GO template library documentation](https://pkg.go.dev/text/template).
