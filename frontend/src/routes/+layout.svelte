<script>
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<link rel="stylesheet" href="/app.css">

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

{@render children?.()}
