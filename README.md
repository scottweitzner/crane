# Crane

#### 🎉 Congratulations, you've found 1000th docker-related project to be named "Crane" 🎉

## What is Crane?

Simply put, Crane is a solution to help templatize your dockerfiles for shared use. Similar to how ansible roles can be centrally stored and consumed -- crane follows the same ethos.

Crane can help team combine common dockerfiles into a single Dockerfile for ingestions from a ci-cd pipeline. 

Crane can also help your stakeholders interact with dockerfiles wiithout having to ever touch a dockerfile.

Why would we want to do that? Let's look at a real world example

![hieracrchy](/docs/assets/hierarchy.jpg)

Let's say you're a research engineer and you manage containerized environment solution for your stakeholders. Each research team has a different of R library requirements. You have a few options:

1. Write a seperate Dockerfile for each team. This would get picked up and used for docker builds in that repositories ci-cd pipeline.

2. Write one Dockerfile with each team's requirements in a separate environement. A user would just have to run `conda activate ...` on startup

3. Create a fancy ci-cd pipeline that builds `n` images on each run based on the specific requirements for each team where `n` is the # of teams.

Ideally you don't want to spend time managing these environments and would rather spend time building the research platform itself, so you let each team manage their own environment. 

What Crane allows you to do is write a single go-templatized Dockerfile in a central git repository and have your downstream teams ingest this within their ci-cd pipelines all without having to touch a Dockerfile.

Example:

**Manifest.yml**
```yaml
source:
  kind: Local
  local: 
    path: /path/to/template.Dockerfile
values:
  installSpecificPackage: false
  copyConfigFile: true
  runArg: my-val

output:
  path: result
  extension: .Dockerfile
```

**template.Dockerfile**
```Dockerfile
FROM python:latest

{{ if .installSpecificPackage }}
RUN apt-get install specificPackage \
    && specificPackage -arg1 -arg2
{{ end }}

{{ if .copyConfigFile }}
COPY configFile /usr/local/etc
{{ end }}

ENTRYPOINT ["./run" ]
CMD ["{{ runArg }}"]
```

This input manifest along with this Dockerfile template will yield the following:

**result.Dockerfile**
```Dockerfile
FROM python:latest

COPY configFile /usr/local/etc

ENTRYPOINT ["./run" ]
CMD ["my-val"]
```

Now I can create a single dockerfile where I can expose different options that will result in different images based on different choices of my end users. 

