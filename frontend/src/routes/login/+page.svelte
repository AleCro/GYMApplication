<script>
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    import { create as createSession } from "$lib/api/session.js";
    import { create as createUser } from "$lib/api/user.js";
    import { cookies } from "$lib/api/cookie.js";
    import { isPasswordSecure } from "$lib/passwordCheck.js";
    import { slide } from "svelte/transition";
    export let data;

    let mode = "login";
    let username = "";
    let password = "";
    let passwordC = "";
    $: passwordCheck = isPasswordSecure(password);
    $: passwordsMatch = password.length > 0 && password === passwordC;

    let message = "";
    let messageType = "";
    let messageVisible = false;
    let messageTimeout;

    function showMessage(msg, type = "success", duration = 4000) {
        clearTimeout(messageTimeout);
        message = msg;
        messageType = type;
        messageVisible = true;
        messageTimeout = setTimeout(() => {
            messageVisible = false;
        }, duration);
    }

    function hideMessage() {
        clearTimeout(messageTimeout);
        messageVisible = false;
        message = "";
        messageType = "";
    }

    async function handleSubmit(event) {
        event.preventDefault();

        // simple client-side validation (already using required on inputs)
        if (!username || !password || (mode === "register" && !username)) {
            showMessage("Please fill out the required fields.", "error");
            return;
        }

        if (mode === "login") {
            createSession(username, password)
                .then((session) => {
                    showMessage(`Successful login, redirecting...`, "success");
                    cookies.setCookie("token", session);
                    goto("/app");
                })
                .catch((err) => {
                    showMessage(err, "error");
                });
        } else {
            if (!passwordsMatch) return;
            if (!passwordCheck.secure) return;
            createUser(username, password).then((session) => {
                    showMessage(`Successful login, redirecting...`, "success");
                    cookies.setCookie("token", session);
                    goto("/app");
                })
                .catch((err) => {
                    showMessage(err, "error");
                });
        }
    }

    function toggleMode() {
        mode = mode === "login" ? "register" : "login";
        hideMessage();
    }
    onMount(() => {
        let urlParams = new URLSearchParams(window.location.search);

        if (urlParams.has("register")) {
            mode = "register";
            urlParams.delete("register");
            let newUrl =
                `${window.location.pathname}?${urlParams.toString()}`.replace(
                    /\?$/,
                    "",
                );
            history.replaceState(null, "", newUrl);
        }
    });
</script>

<svelte:head>
    <title>{mode == "login" ? "Log In" : "Register"} - YSvelGoK</title>
</svelte:head>

<div
    class="bg-gray-900 text-gray-100 min-h-screen flex items-center justify-center p-4 lg:p-8"
>
    <div
        class="w-full max-w-6xl mx-auto h-auto bg-gray-800/70 backdrop-blur-md rounded-2xl shadow-2xl overflow-hidden grid grid-cols-1 lg:grid-cols-5 border border-gray-700/50"
    >
        <div
            class="hidden lg:col-span-2 lg:flex flex-col justify-between p-8 lg:p-12 text-white bg-gray-900/80 border-r border-indigo-500/30"
        >
            <div>
                <div class="inline-flex items-center gap-3 mb-6">
                    <div
                        class="w-10 h-10 rounded-lg bg-gradient-to-br from-[var(--primary-indigo)] to-[var(--primary-pink)] flex items-center justify-center text-white font-bold shadow-lg shrink-0"
                    >
                        YSG
                    </div>
                    <h1 class="text-3xl font-extrabold">YSvelGoK</h1>
                </div>

                <h2 class="text-4xl font-bold leading-tight mb-4">
                    The <span class="gradient-text">full-stack</span> foundation
                    for your next project.
                </h2>
                <p class="text-gray-400 text-lg">
                    Focus on features, not boilerplate. Allowing you to build a
                    robust application using <strong>Svelte</strong>,
                    <strong>Go</strong>
                    and <strong>MongoDB</strong> within minutes.
                </p>
            </div>

            <div class="mt-12 space-y-4">
                <div class="flex items-center space-x-3">
                    <svg
                        class="w-6 h-6 text-indigo-400 shrink-0"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                        aria-hidden="true"
                        ><path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M13 10V3L4 14h7v7l9-11h-7z"
                        ></path></svg
                    >
                    <span class="text-gray-300 font-medium"
                        >Blazing fast API performance.</span
                    >
                </div>
                <div class="flex items-center space-x-3">
                    <svg
                        class="w-6 h-6 text-pink-400 shrink-0"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                        aria-hidden="true"
                        ><path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M9 12l2 2 4-4M17.778 8.086A5.5 5.5 0 0012 5.5v0a5.5 5.5 0 00-5.778 2.586L4.5 10v4c0 1.657 1.343 3 3 3h9c1.657 0 3-1.343 3-3v-4l-2.778-1.914z"
                        ></path></svg
                    >
                    <span class="text-gray-300 font-medium"
                        >Built-in authentication and security.</span
                    >
                </div>
            </div>

            <p class="text-xs text-gray-500 mt-6">
                © {new Date().getFullYear()} YSvelGoK. All rights reserved.
            </p>
        </div>

        <div class="lg:col-span-3 flex items-center justify-center p-6 sm:p-10">
            <div class="w-full max-w-sm">
                <header class="text-center mb-8">
                    <div class="lg:hidden inline-flex items-center gap-3 mb-4">
                        <div
                            class="w-8 h-8 rounded-lg bg-gradient-to-br from-[var(--primary-indigo)] to-[var(--primary-pink)] flex items-center justify-center text-white font-bold shadow-md shrink-0"
                        >
                            YSG
                        </div>
                        <h1 class="text-2xl font-bold">Authorization</h1>
                    </div>
                    <p class="text-gray-400 text-lg">
                        {mode === "login"
                            ? "Sign in to your account"
                            : "Create a new account"}
                    </p>
                </header>
                <form
                    class="space-y-6"
                    on:submit|preventDefault={handleSubmit}
                    aria-describedby="message-box"
                >
                    <div>
                        <label
                            for="email"
                            class="block text-sm font-medium text-gray-300 mb-1"
                            >Username</label
                        >
                        <input
                            id="username"
                            type="text"
                            name="username"
                            bind:value={username}
                            required
                            class="form-input w-full p-3 bg-gray-900 border border-gray-700 rounded-lg text-white placeholder-gray-500 transition duration-150 ease-in-out focus:border-[var(--primary-indigo)]"
                            placeholder="e.g: Cool Guy"
                        />
                    </div>

                    <div>
                        <label
                            for="password"
                            class="block text-sm font-medium text-gray-300 mb-1"
                            >Password</label
                        >
                        <input
                            id="password"
                            type="password"
                            name="password"
                            bind:value={password}
                            required
                            class="form-input w-full p-3 bg-gray-900 border border-gray-700 rounded-lg text-white placeholder-gray-500 transition duration-150 ease-in-out focus:border-[var(--primary-indigo)]"
                            placeholder="••••••••"
                        />
                    </div>
                    {#if password.length > 0 && mode == "register" && !passwordCheck.secure}
                        <div
                            class="p-4 bg-gray-900 border border-gray-700 rounded-md"
                            transition:slide
                        >
                            <ul class="space-y-1 text-sm text-gray-400">
                                {#each passwordCheck.issues as issue}
                                    <li class="flex items-center">
                                        <svg
                                            class="w-4 h-4 mr-2 text-red-400 shrink-0"
                                            fill="none"
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                            stroke-width="2"
                                            viewBox="0 0 24 24"
                                            stroke="currentColor"
                                            ><path d="M6 18L18 6M6 6l12 12"
                                            ></path></svg
                                        >
                                        <span>{issue}</span>
                                    </li>
                                {/each}
                            </ul>
                        </div>
                    {/if}
                    {#if mode === "register"}
                        <div>
                            <label
                                for="confirmPassword"
                                class="block text-sm font-medium text-gray-300 mb-1"
                                >Confirm Password</label
                            >
                            <input
                                id="confirmPassword"
                                type="password"
                                name="confirmPassword"
                                bind:value={passwordC}
                                required
                                class="form-input w-full p-3 bg-gray-900 border border-gray-700 rounded-lg text-white placeholder-gray-500 transition duration-150 ease-in-out focus:border-[var(--primary-indigo)]"
                                placeholder="••••••••"
                            />
                        </div>
                        {#if passwordC.length > 0 && !passwordsMatch}
                            <p
                                class="mt-2 text-sm text-red-400"
                                transition:slide
                            >
                                Passwords do not match.
                            </p>
                        {/if}
                    {/if}

                    <button
                        type="submit"
                        class="form-button w-full py-3 text-lg font-semibold rounded-lg text-white transition duration-150 ease-in-out shadow-lg"
                        class:bg-indigo-600={mode === "login"}
                        class:hover:bg-indigo-700={mode === "login"}
                        class:shadow-indigo-500={mode === "login"}
                        class:bg-[var(--primary-pink)]={mode === "register"}
                        class:hover:bg-pink-700={mode === "register"}
                        class:shadow-pink-500={mode === "register"}
                    >
                        {mode === "login" ? "Log in" : "Register"}
                    </button>
                </form>

                <p class="mt-8 text-center text-gray-400">
                    <span
                        >{mode === "login"
                            ? "Don't have an account?"
                            : "Already have an account?"}</span
                    >
                    <button
                        type="button"
                        class="font-medium gradient-text hover:text-indigo-400 transition duration-150 ml-1"
                        on:click={toggleMode}
                    >
                        {mode === "login" ? "Sign up" : "Log in"}
                    </button>
                </p>

                {#if messageVisible}
                    <div
                        id="message-box"
                        class="mt-4 p-3 rounded-lg text-sm message-enter-leave"
                        role="status"
                        aria-live="polite"
                        class:bg-red-900={messageType === "error"}
                        class:text-red-300={messageType === "error"}
                        class:bg-green-900={messageType === "success"}
                        class:text-green-300={messageType === "success"}
                    >
                        {message}
                    </div>
                {/if}
            </div>
        </div>
    </div>
</div>

<style>
    :root {
        --primary-indigo: #6366f1;
        --primary-pink: #ec4899;
    }

    :global(body) {
        font-family:
            "Inter",
            system-ui,
            -apple-system,
            "Segoe UI",
            Roboto,
            "Helvetica Neue",
            Arial;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
    }

    /* gradient text */
    .gradient-text {
        background-image: linear-gradient(
            90deg,
            var(--primary-indigo),
            var(--primary-pink)
        );
        -webkit-background-clip: text;
        -webkit-text-fill-color: transparent;
        background-clip: text;
        color: transparent;
    }

    /* focus styles for input/button to match the original */
    .form-input:focus,
    .form-button:focus {
        outline: none;
        /* subtle indigo ring */
        box-shadow:
            0 0 0 2px transparent,
            0 0 0 4px rgba(99, 102, 241, 0.18);
    }

    /* small helper for the message box transitions */
    .message-enter-leave {
        transition: all 0.25s ease;
    }
</style>
