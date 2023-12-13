{{ define "content" }}
<p>
    You will be redirected to {{ .Location }} in 3 seconds. If not, <a href="{{ .Location }}">click here</a>.
</p>
<script>
    setTimeout(() => {
        window.location.href = "{{ .Location }}";
    }, 3000);
</script>
{{ end }}
