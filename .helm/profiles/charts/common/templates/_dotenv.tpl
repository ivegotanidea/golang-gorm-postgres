# chart/common/templates/_dotenv.tpl
{{/* Function to generate secret stringData from a .env file */}}
{{- define "common.dotenv.set" -}}
{{- $filePath := .filePath -}}
{{- $files := .Files -}}
{{- range tuple $filePath -}}
{{- $file := $files.Get $filePath | trimSuffix "\n" -}}
    {{- range $line := splitList "\n" $file -}}
    {{- if and (ne $line "") (not (hasPrefix "#" $line)) -}}
        {{- $k := regexFind "^[^=]*" $line -}}
        {{- $v := trimPrefix (printf "%s=" $k) $line -}}
        {{- $k := trim $k -}}
        {{- $v := trim $v -}}
        {{- printf "  %s: \"%s\"" $k (printf "%s" $v) | nindent 0 }}
    {{- end }}
    {{- end -}}
{{- end -}}
{{- end -}}


# chart/common/templates/_dotenv.tpl
{{/* Function to render env variables from a .env file */}}
{{- define "common.dotenv.render" -}}
{{- $filePath := .filePath -}}
{{- $files := .Files -}}
{{- range tuple $filePath -}}
{{- $file := $files.Get $filePath | trimSuffix "\n" -}}
    {{- range $line := splitList "\n" $file -}}
    {{- if and (ne $line "") (not (hasPrefix "#" $line)) -}}
        {{- $k := regexFind "^[^=]*" $line -}}
        {{- $v := trimPrefix (printf "%s=" $k) $line -}}
        {{- $k := trim $k -}}
        {{- $v := trim $v -}}
        {{- $env = set $env $k $v }}
    {{- end }}
    {{- end -}}
{{- end -}}
{{- end -}}
