<script>
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<style>
	* {
		margin: 0px;
	}

	body {
		background-color: #302424;
		color: #f4f4f4;
	}

</style>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

{@render children?.()}
