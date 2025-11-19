<script>
    import "../app.css";
    import { page } from '$app/stores';
    
    $: status = $page.status || 500;
    $: message = $page.error?.message || 'An unexpected error occurred.';
    $: title = (status === 404) ? 'Page Not Found' : 'Server Error';
    $: path = $page.url.pathname;
</script>

<svelte:head>
    <title>{status} {title} | Starter API</title>
</svelte:head>

<div class="flex flex-col items-center justify-center text-center p-8 space-y-6 min-h-screen">
    <h1 class="text-7xl font-extrabold text-indigo-400 tracking-tight">{status}</h1>
    <h2 class="text-4xl font-bold text-white">{title}</h2>
    <p class="text-xl text-gray-400 max-w-lg">
        {message}
    </p>

    <p class="text-sm text-gray-500">
        Attempted path: <code class="font-mono">{path}</code>
    </p>

    <div class="flex flex-col sm:flex-row justify-center gap-4 pt-4">
        <a 
            href="/app"
            class="px-6 py-3 bg-indigo-600 text-white font-semibold rounded-lg shadow-lg hover:bg-indigo-700 transition duration-150 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-indigo-500 focus-visible:ring-offset-2 focus-visible:ring-offset-gray-900"
        >
            Go to App
        </a>
        <a 
            href="/"
            class="px-6 py-3 border border-gray-600 text-gray-300 font-semibold rounded-lg hover:bg-gray-800 transition duration-150 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-gray-400 focus-visible:ring-offset-2 focus-visible:ring-offset-gray-900"
        >
            Return Home
        </a>
    </div>
</div>
