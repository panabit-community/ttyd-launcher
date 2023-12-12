{{ define "content" }}
<button onclick="clickHandler()">Launch ttyd</button>
<script>
    function clickHandler() {
        setTimeout(() => {
            window.location = "{{ .WindowLocation }}";
        }, 1000);
    }
</script>
{{ end }}
