<script>
  import favicon from '$lib/assets/favicon.svg';
</script>

<link rel="stylesheet" href="/app.css">

<svelte:head>
  <link rel="icon" href={favicon} />
  <title>GYM App</title>
</svelte:head>

<nav style="margin-bottom: 20px;">
  <a href="/login" style="margin-right: 10px;">Login</a>
  <a href="/notes" style="margin-right: 10px;">Notes</a>
  <a href="/calendar" style="margin-right: 10px;">Calendar</a>
  <a href="/exercise" style="margin-right: 10px;">Exercise</a>
</nav>

<slot /> <!-- render current page -->
