<script>
  import { onMount } from 'svelte';
  // ...
</script>
<div class="progress-layout">
  <slot />
</div>

<style>
  .progress-layout {
    padding: 1rem;
  }
</style>