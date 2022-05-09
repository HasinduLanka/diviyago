<script lang="ts">
  import type { ConvertedResult } from "./APITypes";

  let msg = "";

  let endpoint = "/api/imageconvert";
  let contentType = "";

  let fileSize = "";

  let result: ConvertedResult | null;

  $: code = `
    fetch(${document.location.origin}${endpoint}, {
      method: "POST",
      headers: [
        ["Content-Type", "${contentType}"],
      ],
      body: file,
    })
    `;

  let init = async () => {};

  init();

  let filelst: FileList;

  function fileChanged() {
    if (filelst && filelst.length > 0) {
      const file = filelst[0];
      fileSize = (file.size / 1024).toFixed(0) + " KB";
      contentType = file.type;
    }
  }

  function submitForm() {
    event.preventDefault();

    if (filelst && filelst.length > 0) {
      const file = filelst[0];

      fetch(endpoint, {
        method: "POST",
        headers: [["Content-Type", contentType]],
        body: file,
      })
        .then(async function (response) {
          let resp = await response.json();
          msg = JSON.stringify(resp, null, 2);

          try {
            result = resp as ConvertedResult;
          } catch (error) {
            result = null;
            msg = await response.text();
            return;
          }

          if (result.Success) {
            console.log("Successfully converted");
          } else {
            result = null;
            msg = JSON.stringify(resp);
          }
        })
        .catch(async (error) => {
          
          msg = await error;
        });
    }
  }
</script>

<main>
  <div style="margin: 4rem;">
    <h1>DiviyaGo</h1>
    <p>Self contained image processing library for Go</p>
  </div>

  <div>
    <form on:submit={submitForm}>
      <div class="box">
        <span>Endpoint:</span>
        <input type="text" bind:value={endpoint} />
        <br />
      </div>

      <div class="box">
        <input type="file" bind:files={filelst} on:change={fileChanged} />
        <p>{fileSize}</p>
        <br />
      </div>

      <div class="box">
        <span>Content Type:</span>
        <input type="text" bind:value={contentType} placeholder={"image/png"} />
      </div>
      <div class="box">
        <input type="submit" />
      </div>
    </form>

    <div class="box">
      <p>Fetch call:</p>
      <pre> {code} </pre>
      <br />
      <br />
    </div>

    <div class="box">
      <p>Result preview:</p>
      {#if result}
        {#each Object.keys(result.TransformedResults) as transformedResultID}
          <div class="block">
            <h4>{transformedResultID}</h4>
            <p>
              Scale: {JSON.stringify(
                result.TransformedResults[transformedResultID].Scale
              )}
            </p>
            <p>
              Codec: {result.TransformedResults[transformedResultID].VideoCodec}
            </p>
            <p>
              Size: {result.TransformedResults[transformedResultID].Data
                .length / 1024} KB
            </p>
            <img
              src={`data:${result.TransformedResults[transformedResultID].VideoCodec};base64,${result.TransformedResults[transformedResultID].Data}`}
              alt={transformedResultID}
            />
          </div>
        {/each}
      {/if}

      <p>Response:</p>
      <pre> {msg} </pre>
      <br />
      <br />
    </div>
  </div>
</main>

<style>
  .box {
    padding: 2vw;
    margin: 2vw;
    box-shadow: 0 0 1rem rgba(0, 0, 0, 0.5);
    width: 90vw;
    overflow-x: auto;
  }

  .block {
    padding: 1vw;
    margin-left: auto;
    margin-right: auto;
    box-shadow: 0 0 0.5rem rgba(0, 0, 0, 0.5);
    width: 95%;
    overflow-x: auto;
  }

  pre {
    background-color: gray;
    width: 100%;
    padding: 1rem;
  }
</style>
