<script>
    import "../../app.css";
    import { onDestroy, onMount } from "svelte";
    import { refresh as refreshSession } from "$lib/api/session";
    import { cookies } from "$lib/api/cookie";
    import { userStore } from "$lib/api/stores.js";
    import { writable } from "svelte/store";
    import { quintOut } from "svelte/easing";
    import { slide } from "svelte/transition";
    // User reactivity on session update
    export let data;
    userStore.set(data?.user ?? null);
    $: user = data.user;
    let unsubscribe = userStore.subscribe((value) => {
        user = value;
    });

    let sessionRefreshInterval = null;

    let mobileOpen = writable(false);

    onMount(() => {
        if (sessionRefreshInterval != null) {
            clearInterval(sessionRefreshInterval);
        }
        sessionRefreshInterval = setInterval(() => {
            let expiresInMS = user.exp * 1000 - new Date().valueOf();
            if (expiresInMS < 1000 * 60 * 5) {
                refreshSession(cookies.token)
                    .then((res) => {
                        cookies.token = res;
                        fetch("/app", {
                            method: "POST",
                            headers: { Authorization: res },
                        })
                            .then((res) => res.json())
                            .then((res) => {
                                userStore.set(res);
                                data.user = res;
                            })
                            .catch((err) => {
                                console.error("Unable to renew user data", err);
                            });
                    })
                    .catch((err) => {
                        console.error("Unable to renew session", err);
                    });
            }
        }, 1000*60);
    });

    onDestroy(() => {
        clearInterval(sessionRefreshInterval);
        sessionRefreshInterval = null;
        unsubscribe();
    });

    let mainNav = [
        { href: "/app", label: "Dashboard", icon: "🏠" },
        { href: "/app/notes", label: "Notes", icon: "📝" },
        { href: "/app/calendar", label: "Calendar", icon: "📅" },
        { href: "/app/workouts", label: "Workouts", icon: "💪" },
        { href: "/app/progress", label: "Progress", icon: "📈" },
        {
            href: "/app/user-management",
            label: "User Management",
            icon: "👤",
            group: 255,
        },
    ];

    let secondaryNav = [
        { href: "/app/documentation", label: "Documentation", icon: "📚" },
        { href: "/logout", label: "Sign out", icon: "🚪" },
    ];

    async function signOut() {
        window.location.href = "/logout";
    }

    $: userInitial = user
        ? user.username
            ? user.username.charAt(0).toUpperCase()
            : "G"
        : "U";
</script>

<svelte:head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width,initial-scale=1" />
    <title>{data?.title ?? "AleGYM"}</title>
    <meta name="description" content="AleGYM — Your Ultimate Fitness Companion" />
    <link rel="icon" href="/favicon.ico" />
</svelte:head>

    <div class="min-h-screen bg-background text-text font-sans flex flex-col bg-[url('https://images.unsplash.com/photo-1534438327276-14e5300c3a48?q=80&w=2070&auto=format&fit=crop')] bg-cover bg-center bg-fixed bg-no-repeat">
        <div class="absolute inset-0 bg-black/60 backdrop-blur-[2px] z-0"></div>
        
        <header
            class="glass sticky top-0 z-50 border-b-0"
        >
            <div class="max-w-screen-2xl mx-auto px-4 sm:px-6 lg:px-8">
                <div class="flex items-center justify-between h-16">
                    <div class="flex items-center gap-4">
                        <button
                            class="p-2 rounded-md hover:bg-white/10 transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary lg:hidden"
                            aria-label="Toggle sidebar"
                            on:click={() => mobileOpen.update((v) => !v)}
                        >
                            {#if $mobileOpen}
                                <svg
                                    class="w-6 h-6"
                                    viewBox="0 0 24 24"
                                    fill="none"
                                    stroke="currentColor"
                                    ><path
                                        stroke-width="1.5"
                                        stroke-linecap="round"
                                        stroke-linejoin="round"
                                        d="M6 18L18 6M6 6l12 12"
                                    /></svg
                                >
                            {:else}
                                <svg
                                    class="w-6 h-6"
                                    viewBox="0 0 24 24"
                                    fill="none"
                                    stroke="currentColor"
                                    ><path
                                        stroke-width="1.5"
                                        stroke-linecap="round"
                                        stroke-linejoin="round"
                                        d="M4 7h16M4 12h16M4 17h16"
                                    /></svg
                                >
                            {/if}
                        </button>
    
                        <a
                            href="/app"
                            class="flex items-center gap-3 focus:outline-none focus-visible:ring-2 focus-visible:ring-primary rounded-lg -m-1 p-1"
                        >
                            <div
                                class="w-9 h-9 rounded-lg bg-gradient-to-br from-primary to-secondary flex items-center justify-center text-white font-bold shadow-md shrink-0"
                            >
                                AG
                            </div>
                            <div class="hidden sm:block">
                                <div class="font-bold text-lg tracking-wide">AleGYM</div>
                            </div>
                        </a>
                    </div>
    
                    <div class="flex items-center gap-4">
                        <div
                            class="hidden md:flex items-center glass-input rounded-lg px-3 py-1.5 focus-within:ring-2 focus-within:ring-primary transition"
                        >
                            <svg
                                class="w-5 h-5 mr-2 text-gray-400"
                                viewBox="0 0 24 24"
                                fill="none"
                                aria-hidden
                                stroke="currentColor"
                            >
                                <path
                                    stroke-width="1.5"
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    d="M21 21l-4.35-4.35M11 18a7 7 0 100-14 7 7 0 000 14z"
                                />
                            </svg>
                            <input
                                placeholder="Search..."
                                class="bg-transparent outline-none text-sm placeholder-gray-400 w-40 text-text"
                            />
                        </div>
    
                        <a
                            href="/app/me"
                            class="flex items-center gap-2 p-1.5 rounded-full hover:bg-white/10 transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary"
                        >
                            <div
                                class="w-8 h-8 rounded-full bg-secondary text-white flex items-center justify-center text-sm font-medium shadow-md shrink-0"
                            >
                                {userInitial}
                            </div>
                            <span
                                class="hidden lg:block text-sm font-medium text-text mr-1" title="Session expires: {user.exp - new Date().valueOf()/1000}m"
                                >{data?.user?.username ?? "Guest"}</span
                            >
                        </a>
    
                        <button
                            on:click={signOut}
                            class="hidden sm:block px-3 py-2 rounded-lg bg-accent/20 text-accent text-sm font-medium hover:bg-accent/30 transition shadow-sm focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent"
                        >
                            Sign out
                        </button>
                    </div>
                </div>
            </div>
        </header>
    
        <div class="flex-grow flex max-w-screen-2xl mx-auto w-full min-w-0 z-10 relative">
            <aside
                class="hidden lg:block w-64 flex-shrink-0 glass border-y-0 border-l-0 border-r-0 pt-6 px-4 pb-4 sticky top-16 h-[calc(100vh-4rem)] overflow-hidden ml-4 my-4 rounded-xl"
            >
                <nav class="space-y-6">
                    <div class="space-y-1">
                        <div
                            class="text-xs uppercase tracking-wider text-gray-400 font-semibold mb-2"
                        >
                            Workspace
                        </div>
                        <ul class="space-y-1 text-sm">
                            {#each mainNav as item}
                                {#if !item.group || item.group == data.user?.group}
                                    <li>
                                        <a
                                            href={item.href}
                                            class="flex items-center gap-3 px-3 py-2 rounded-lg hover:bg-white/10 transition font-medium text-gray-300 hover:text-primary focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary"
                                        >
                                            <span
                                                class="text-lg w-5 flex justify-center"
                                                >{item.icon}</span
                                            >
                                            {item.label}
                                        </a>
                                    </li>
                                {/if}
                            {/each}
                        </ul>
                    </div>
    
                    <div class="space-y-1 pt-4 border-t border-white/10">
                        <div
                            class="text-xs uppercase tracking-wider text-gray-400 font-semibold mb-2"
                        >
                            General
                        </div>
                        <ul class="space-y-1 text-sm">
                            {#each secondaryNav as item}
                                {#if !item.group || item.group == data.user?.group}
                                    <li>
                                        <a
                                            href={item.href}
                                            class="flex items-center gap-3 px-3 py-2 rounded-lg hover:bg-white/10 transition font-medium text-gray-300 hover:text-primary focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary"
                                        >
                                            <span
                                                class="text-lg w-5 flex justify-center"
                                                >{item.icon}</span
                                            >
                                            {item.label}
                                        </a>
                                    </li>
                                {/if}
                            {/each}
                        </ul>
                    </div>
                </nav>
            </aside>
    
            {#if $mobileOpen}
                <div
                    class="lg:hidden fixed inset-0 z-40 bg-black/70 transition-opacity duration-300 ease-in-out"
                    on:click={() => mobileOpen.set(false)}
                ></div>
    
                <aside
                    id="mobile-sidebar"
                    transition:slide={{ duration: 250, easing: quintOut }}
                    class="lg:hidden fixed left-0 top-0 bottom-0 z-50 w-72 p-6 glass shadow-2xl overflow-y-auto"
                >
                    <div class="flex items-center justify-between mb-8">
                        <a href="/" class="flex items-center gap-3">
                            <div
                                class="w-9 h-9 rounded-lg bg-gradient-to-br from-primary to-secondary flex items-center justify-center text-white font-bold"
                            >
                                AG
                            </div>
                            <div class="text-base font-bold">AleGYM</div>
                        </a>
                        <button
                            on:click={() => mobileOpen.set(false)}
                            aria-label="Close drawer"
                            class="p-2 rounded-md hover:bg-white/10 transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary"
                        >
                            <svg
                                class="w-6 h-6"
                                viewBox="0 0 24 24"
                                fill="none"
                                stroke="currentColor"
                                ><path
                                    stroke-width="1.5"
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    d="M6 18L18 6M6 6l12 12"
                                /></svg
                            >
                        </button>
                    </div>
    
                    <nav class="space-y-6">
                        <div class="space-y-1">
                            <div
                                class="text-xs uppercase tracking-wider text-gray-400 font-semibold mb-2"
                            >
                                Workspace
                            </div>
                            <ul class="space-y-1 text-base">
                                {#each [...mainNav, ...secondaryNav] as item}
                                    {#if !item.group || item.group == data.user?.group}
                                        <li>
                                            <a
                                                href={item.href}
                                                on:click={() =>
                                                    mobileOpen.set(false)}
                                                class="flex items-center gap-3 px-3 py-2 rounded-lg hover:bg-white/10 transition font-medium text-gray-300 hover:text-primary focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary"
                                            >
                                                <span
                                                    class="text-xl w-5 flex justify-center"
                                                    >{item.icon}</span
                                                >
                                                {item.label}
                                            </a>
                                        </li>
                                    {/if}
                                {/each}
                            </ul>
                        </div>
    
                        <div class="mt-10 pt-4 border-t border-white/10">
                            <a
                                href="/me"
                                on:click={() => mobileOpen.set(false)}
                                class="flex items-center gap-3 px-2 py-2 rounded-lg hover:bg-white/10 transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary"
                            >
                                <div
                                    class="w-8 h-8 rounded-full bg-secondary text-white flex items-center justify-center text-sm font-medium shrink-0"
                                >
                                    {userInitial}
                                </div>
                                <div>
                                    <div class="font-medium">
                                        {data?.user?.username ?? "Guest"}
                                    </div>
                                    <div class="text-xs text-gray-400">
                                        View profile
                                    </div>
                                </div>
                            </a>
                            <button
                                on:click={signOut}
                                class="mt-4 inline-block w-full text-center px-3 py-2 rounded-lg bg-accent/20 text-accent text-sm font-medium hover:bg-accent/30 transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent"
                                >Sign out</button
                            >
                        </div>
                    </nav>
                </aside>
            {/if}
    
            <main class="flex-grow p-4 sm:p-6 lg:p-8 min-w-0">
                <div
                    class="p-6 rounded-xl glass transition shadow-xl w-full max-w-full overflow-hidden"
                >
                    <slot />
                </div>
            </main>
        </div>
    
        <footer class="border-t border-white/10 glass mt-auto">
            <div
                class="max-w-screen-2xl mx-auto px-4 sm:px-6 lg:px-8 py-4 text-sm text-gray-400"
            >
                <div class="flex items-center justify-between">
                    <div>© {new Date().getFullYear()} AleGYM</div>
                    <div class="flex items-center gap-4">
                        <a
                            href="/privacy"
                            class="hover:text-primary transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary rounded-sm"
                            >Privacy</a
                        >
                    </div>
                </div>
            </div>
        </footer>
    </div>
