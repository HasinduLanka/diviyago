<script lang="ts">
  let msg;

  let init = async () => {
    let helloResp = await fetch("/api/hello");
    msg = await helloResp.text();
  };

  init();
</script>

<main style="margin: 4rem; text-align: center;">
  <h1>DiviyaGo</h1>
  <p>Platform independent image manipulation library for Go</p>
  <pre> {msg} </pre>
</main>
