<script>
	import favicon from '$lib/assets/favicon.svg';
	import "../app.css";

	let { children } = $props();
</script>

<link rel="stylesheet" href="/assets/app.css">

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

{@render children?.()}
