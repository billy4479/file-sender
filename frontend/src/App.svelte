<script lang="ts">
  import './tailwind.css';
  import './styles.css';

  import Send from './Send.svelte';
  import Recive from './Recive.svelte';

</script>

<h1 class="text-5xl text-red-500 mb-7">File sender</h1>

<div class="flex-container">
  <div>
    <Send />
  </div>
  <span>Or</span>
  <div>
    <Recive />
  </div>
</div>

<style global type="postcss">
  body {
    width: 85%;
    margin: 5% auto;
    text-align: center;
  }

  .flex-container {
    display: flex;
    flex-direction: row;
    justify-content: space-evenly;
    align-items: center;
  }

  .flex-container > div {
    width: 40%;
    padding: 1em;
    overflow: hidden;
    @apply rounded shadow;
  }

  h2 {
    @apply text-3xl text-green-600;
  }

  @media screen and (max-width: 480px) {
    .flex-container {
      flex-direction: column;
    }
    .flex-container > div {
      margin-top: 1rem;
      margin-bottom: 1rem;
      width: 100%;
    }
  }

</style>
