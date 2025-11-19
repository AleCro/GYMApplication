<script>
    import { request } from "$lib/api/util";
    import { cookies } from "$lib/api/cookie.js";
    import { quintOut } from "svelte/easing";
    import { fly } from "svelte/transition";

    import { userStore } from "$lib/api/stores.js";
    import { onDestroy } from "svelte";
    export let data;
    userStore.set(data?.user ?? null);
    $: user = data.user;
    let unsubscribe = userStore.subscribe((value) => {
        user = value;
    });
    onDestroy(() => {
        unsubscribe();
    });
    
    let users = data.users || [];
    let total = data.total || 0;
    let limit = data.limit || 10;
    let currentPage = data.page || 1;
    let isLoading = false;
    let error = null;

    let searchTerm = "";
    let searchTimeout;
    let isSearching = false;

    let isEditing = false;
    let editingUser = null;
    let newUsername = "";
    let newRole = 0;

    const ROLE_MAP = {
        0: "User",
        255: "Administrator",
    };

    const getRoleName = (role) => {
        return ROLE_MAP[role] || `Custom (${role})`;
    };

    async function fetchUsers(page, search = "") {
        isLoading = true;
        error = null;
        currentPage = page;

        try {
            let path = `/users?limit=${limit}&page=${page}`;
            if (search.trim()) {
                path += `&search=${encodeURIComponent(search.trim())}`;
            }

            const response = await request(path, "GET", null, {
                Authorization: `Bearer ${cookies.token}`,
            });

            users = response.users || [];
            total = response.total || 0;
            limit = response.limit || 10;
            currentPage = response.page || 1;
        } catch (err) {
            error = "Failed to load users: " + (err.message || "Network error");
            console.error(err);
        } finally {
            isLoading = false;
            isSearching = false;
        }
    }

    function handleSearch() {
        isSearching = true;

        if (searchTimeout) {
            clearTimeout(searchTimeout);
        }

        searchTimeout = setTimeout(() => {
            fetchUsers(1, searchTerm);
        }, 300);
    }

    function clearSearch() {
        searchTerm = "";
        fetchUsers(1);
    }

    function startEdit(user) {
        editingUser = user;
        newUsername = user.username;
        newRole = user.group;
        isEditing = true;
    }

    async function saveUser() {
        if (!editingUser) return;

        const path = `/users/${editingUser.id}`;
        const body = JSON.stringify({
            username: newUsername,
            group: parseInt(newRole, 10),
        });

        try {
            let response = await request(path, "PUT", body, {
                Authorization: `Bearer ${cookies.token}`,
            });

            users = users.map((u) => (u.id === editingUser.id ? response : u));

            isEditing = false;
            editingUser = null;
        } catch (err) {
            alert(
                "Failed to update user: " +
                    (err.message || "Check server logs."),
            );
            console.error("Update failed:", err);
        }
    }

    $: {
        if (users.length === 0 && !isLoading) {
            fetchUsers(currentPage);
        }
    }
</script>

<svelte:head>
    <title>User Management</title>
</svelte:head>

<div class="space-y-6">
    <header class="pb-4 border-b border-gray-700">
        <h1 class="text-3xl font-bold text-white">👤 User Management</h1>
        <p class="mt-1 text-gray-400">
            Manage user accounts, roles, and details. Total Users: {total}
        </p>
    </header>

    <!-- Search Bar -->
    <div class="bg-gray-800 p-4 rounded-lg border border-gray-700">
        <div class="flex gap-4 items-center">
            <div class="flex-1 relative">
                <div
                    class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
                >
                    <svg
                        class="h-5 w-5 text-gray-400"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                        />
                    </svg>
                </div>
                <input
                    type="text"
                    bind:value={searchTerm}
                    on:input={handleSearch}
                    placeholder="Search by username or ID..."
                    class="w-full pl-10 pr-4 py-2 bg-gray-700 border border-gray-600 rounded-lg text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
                />
                {#if searchTerm}
                    <button
                        on:click={clearSearch}
                        class="absolute inset-y-0 right-0 pr-3 flex items-center text-gray-400 hover:text-white"
                    >
                        <svg
                            class="h-5 w-5"
                            fill="none"
                            viewBox="0 0 24 24"
                            stroke="currentColor"
                        >
                            <path
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="2"
                                d="M6 18L18 6M6 6l12 12"
                            />
                        </svg>
                    </button>
                {/if}
            </div>
            {#if searchTerm}
                <button
                    on:click={clearSearch}
                    class="px-4 py-2 text-sm font-medium text-gray-300 bg-gray-700 hover:bg-gray-600 rounded-lg transition"
                >
                    Clear
                </button>
            {/if}
        </div>
        {#if searchTerm}
            <p class="mt-2 text-sm text-gray-400">
                Searching for: "{searchTerm}"
                {#if isSearching}
                    <span class="ml-2">Searching...</span>
                {/if}
            </p>
        {/if}
    </div>

    {#if error}
        <div
            class="p-4 bg-red-900/30 text-red-300 rounded-lg border border-red-700 font-medium"
            role="alert"
        >
            {error}
        </div>
    {/if}

    <div
        class="overflow-x-auto rounded-lg border border-gray-700 shadow-xl bg-gray-800"
    >
        <table class="min-w-full divide-y divide-gray-700">
            <thead class="bg-gray-700/50">
                <tr>
                    <th
                        scope="col"
                        class="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider"
                        >ID</th
                    >
                    <th
                        scope="col"
                        class="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider"
                        >Username</th
                    >
                    <th
                        scope="col"
                        class="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider"
                        >Role</th
                    >
                    <th scope="col" class="relative px-6 py-3"
                        ><span class="sr-only">Edit</span></th
                    >
                </tr>
            </thead>
            <tbody class="divide-y divide-gray-800">
                {#if isLoading}
                    <tr>
                        <td
                            colspan="5"
                            class="px-6 py-4 text-center text-gray-500 italic"
                        >
                            {#if isSearching}
                                Searching users...
                            {:else}
                                Loading users...
                            {/if}
                        </td>
                    </tr>
                {:else if users.length === 0}
                    <tr>
                        <td
                            colspan="5"
                            class="px-6 py-4 text-center text-gray-500 italic"
                        >
                            {#if searchTerm}
                                No users found matching "{searchTerm}".
                            {:else}
                                No users found.
                            {/if}
                        </td>
                    </tr>
                {:else}
                    {#each users as user (user.id)}
                        <tr
                            class="transition duration-150 {user.id ==
                            data?.user.userID
                                ? 'bg-gray-700 hover:bg-gray-900'
                                : 'hover:bg-gray-700'}"
                        >
                            <td
                                class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-300"
                                >{user?.id}</td
                            >
                            <td
                                class="px-6 py-4 whitespace-nowrap text-sm text-white font-medium"
                            >
                                {user.username}
                            </td>
                            <td
                                class="px-6 py-4 whitespace-nowrap text-sm text-gray-400"
                            >
                                <span
                                    class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full"
                                    class:bg-indigo-600={user.group >= 100}
                                    class:bg-blue-600={user.group < 100}
                                >
                                    {getRoleName(user.group)}
                                </span>
                            </td>
                            <td
                                class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium"
                            >
                                <button
                                    on:click={() => startEdit(user)}
                                    class="text-indigo-400 hover:text-indigo-300"
                                >
                                    Edit
                                </button>
                            </td>
                        </tr>
                    {/each}
                {/if}
            </tbody>
        </table>
    </div>

    <div class="flex justify-between items-center pt-4">
        <p class="text-sm text-gray-400">
            Showing {Math.min(limit * (currentPage - 1) + 1, total)} to {Math.min(
                limit * currentPage,
                total,
            )} of {total} results.
            {#if searchTerm}
                <span class="block text-xs text-gray-500 mt-1"
                    >Filtered by: "{searchTerm}"</span
                >
            {/if}
        </p>
        <div class="flex gap-2">
            <button
                on:click={() => fetchUsers(currentPage - 1, searchTerm)}
                disabled={currentPage === 1 || isLoading}
                class="px-4 py-2 text-sm font-medium rounded-md transition disabled:opacity-50 disabled:cursor-not-allowed"
                class:bg-gray-700={currentPage !== 1}
                class:hover:bg-gray-600={currentPage !== 1}
                class:text-white={currentPage !== 1}
            >
                Previous
            </button>
            <button
                on:click={() => fetchUsers(currentPage + 1, searchTerm)}
                disabled={currentPage * limit >= total || isLoading}
                class="px-4 py-2 text-sm font-medium rounded-md transition disabled:opacity-50 disabled:cursor-not-allowed"
                class:bg-gray-700={currentPage * limit < total}
                class:hover:bg-gray-600={currentPage * limit < total}
                class:text-white={currentPage * limit < total}
            >
                Next
            </button>
        </div>
    </div>
</div>

{#if isEditing}
    <div
        transition:fly={{ y: -10, duration: 250, easing: quintOut }}
        class="fixed inset-0 bg-black/70 z-[100] flex items-center justify-center p-4"
        on:click|self={() => (isEditing = false)}
    >
        <div
            class="bg-gray-800 p-8 rounded-xl shadow-2xl w-full max-w-md border border-gray-700"
        >
            <h3 class="text-xl font-bold text-white mb-4">
                Edit User: {editingUser.username}
            </h3>

            <div class="space-y-4">
                <div>
                    <label
                        for="edit-username"
                        class="block text-sm font-medium text-gray-300"
                        >Username (String)</label
                    >
                    <input
                        id="edit-username"
                        type="text"
                        bind:value={newUsername}
                        class="mt-1 block w-full px-3 py-2 bg-gray-700 border border-gray-600 rounded-md text-white shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    />
                </div>

                <div>
                    <label
                        for="edit-role"
                        class="block text-sm font-medium text-gray-300"
                        >Role (Number 0-255)</label
                    >
                    <input
                        id="edit-role"
                        type="number"
                        min="0"
                        max="255"
                        bind:value={newRole}
                        class="mt-1 block w-full px-3 py-2 bg-gray-700 border border-gray-600 rounded-md text-white shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    />
                    <p class="mt-1 text-xs text-gray-500">
                        {getRoleName(editingUser.group)}
                    </p>
                </div>
            </div>

            <div class="mt-6 flex justify-end gap-3">
                <button
                    on:click={() => (isEditing = false)}
                    class="px-4 py-2 text-sm font-medium rounded-lg text-gray-300 bg-gray-700 hover:bg-gray-600 transition"
                >
                    Cancel
                </button>
                <button
                    on:click={saveUser}
                    class="px-4 py-2 text-sm font-medium rounded-lg text-white bg-indigo-600 hover:bg-indigo-700 transition"
                    disabled={isLoading}
                >
                    Save Changes
                </button>
            </div>
        </div>
    </div>
{/if}
