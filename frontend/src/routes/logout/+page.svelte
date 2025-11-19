<script>
    import { onMount } from "svelte";
    import { goto } from '$app/navigation';
    import { remove as removeSession } from '$lib/api/session';
    import { cookies } from '$lib/api/cookie';

    onMount(() => {
        removeSession(`Bearer ${cookies.token}`).then(() => {
            delete cookies.token;
            window.location = "/login";
        }).catch(err => {
            delete cookies.token;
            window.location = "/login";
        });
    });
</script>

<svelte:head>
    <title>Logging Out...</title>
    <meta name="robots" content="noindex, nofollow" />
</svelte:head>

<div
    class="min-h-screen flex flex-col items-center justify-center bg-gray-900 text-gray-100 font-sans relative overflow-hidden"
>
    <div
        class="absolute inset-0 bg-[radial-gradient(ellipse_at_center,_var(--tw-gradient-stops))] from-indigo-900/20 via-gray-900 to-gray-900"
    ></div>

    <div
        class="relative z-10 space-y-8 text-center animate-fadeIn max-w-md px-4"
    >
        <div class="flex justify-center">
            <svg
                class="w-8 h-8 text-indigo-400 animate-pulse"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
            >
                <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
                ></path>
            </svg>
        </div>

        <div class="space-y-3">
            <h1
                class="text-3xl sm:text-4xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-indigo-200 tracking-tight"
            >
                Logging you out...
            </h1>
            <p class="text-lg text-gray-400 font-medium">
                Wrapping things up. See you next time! 👋
            </p>
        </div>

        <div
            class="w-64 h-1.5 bg-gray-800 rounded-full overflow-hidden mx-auto relative"
        >
            <div
                class="absolute top-0 bottom-0 left-0 w-1/3 bg-indigo-500 rounded-full animate-indeterminate"
            ></div>
        </div>
    </div>
</div>

<style>
    /* Smoother fade in for the main content */
    @keyframes fadeIn {
        from {
            opacity: 0;
            transform: translateY(20px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }

    .animate-fadeIn {
        animation: fadeIn 0.7s cubic-bezier(0.4, 0, 0.2, 1) forwards;
    }

    /* Standard indeterminate sliding animation */
    @keyframes indeterminate {
        0% {
            left: -35%;
            right: 100%;
        }
        60%,
        100% {
            left: 100%;
            right: -90%;
        }
    }

    .animate-indeterminate {
        /* w-1/3 is set in tailwind class above, this animates it across */
        animation: indeterminate 2s cubic-bezier(0.65, 0.815, 0.735, 0.395)
            infinite;
    }

    /* Respect user reduced motion settings */
    @media (prefers-reduced-motion: reduce) {
        .animate-fadeIn,
        .animate-spin,
        .animate-indeterminate {
            animation: none !important;
            opacity: 1 !important;
            transform: none !important;
        }
    }
</style>
