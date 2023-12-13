{{ define "content" }}
<button onclick="clickHandler()">Launch ttyd</button>
<script>
    function clickHandler() {
        window.location.href = "/cgi-bin/App/ttyd/api";
    }
</script>
{{ end }}
