{{- define "webhook-receiver.name" -}}
webhook-receiver
{{- end }}

{{- define "webhook-receiver.fullname" -}}
{{ printf "%s-%s" .Release.Name (include "webhook-receiver.name" .) }}
{{- end }}
