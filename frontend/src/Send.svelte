<script lang="ts">
  let files: FileList;
  let uploading = false;
  let errorMsg = '';
  let id = '';
  let status = '';

  async function onSubmit() {
    if (files.length === 0) return;
    uploading = true;
    errorMsg = '';
    id = await fetch(`http://${window.location.host}/upload`)
      .then((data) => data.text())
      .then((text) => text);

    const formData = new FormData();
    formData.append('file', files[0]);

    fetch(`http://${window.location.host}/upload/${id}`, {
      method: 'POST',
      body: formData,
    })
      .then((res) => {
        if (!res.ok) {
          res.text().then((err) => {
            errorMsg = err;
          });
        }
      })
      .catch((err) => (errorMsg = err))
      .finally(() => {
        uploading = false;
      });

    setTimeout(() => {
      const ws = new WebSocket(`ws://${window.location.host}/status?id=${id}`);
      ws.onmessage = (e: MessageEvent<any>) => {
        status = e.data;
      };
      ws.onerror = () => {
        errorMsg = 'An error has occurred';
      };
    }, 100);
  }

</script>

<h2>Send a file</h2>
<input class="m-4" type="file" bind:files />
<br />
{#if !uploading}
  <button class="rounded shadow px-3 py-2" on:click|preventDefault={onSubmit}
    >Send</button
  >
{:else}
  <p>
    Use this information to download the file. Every link can be used only one
    time.
  </p>
  <br />
  <span>Id:</span>
  <pre>{id}</pre>
  <br />
  <span>Link:</span>
  <a
    class="underline text-blue-700"
    href={`http://${window.location.host}/download/${id}`}
    >{`http://${window.location.host}/download/${id}`}</a
  >
  <p>{status}</p>
{/if}
{#if errorMsg !== ''}
  <pre> {errorMsg} </pre>
{/if}

<style>
</style>
