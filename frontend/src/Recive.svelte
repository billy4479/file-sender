<script lang="ts">
  let id = '';

  function onClick() {
    window.open(`http://${window.location.host}/download/${id}`);
  }

</script>

<h2>Recive a file</h2>

<input
  type="text"
  bind:value={id}
  class="shadow rounded p-3 m-4"
  placeholder="Paste the file ID here"
/>

<button class="rounded shadow px-3 py-2" on:click={onClick}> Download </button>

<style>
</style>
