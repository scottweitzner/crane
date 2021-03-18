FROM busybox:latest

{{ if .sayHello -}}
RUN echo "hello"
{{- end }}


