<script>
    import { onMount } from "svelte";
    import { cookies } from "$lib/api/cookie";
    import { request } from "$lib/api/util";

    let events = [];
    let currentDate = new Date();
    let selectedDate = null;
    let showModal = false;

    // Modal form data
    let eventTitle = "";
    let eventDesc = "";
    let editingEventId = null;

    $: year = currentDate.getFullYear();
    $: month = currentDate.getMonth();
    $: daysInMonth = new Date(year, month + 1, 0).getDate();
    $: firstDayOfMonth = new Date(year, month, 1).getDay();
    $: monthName = currentDate.toLocaleString("default", { month: "long" });

    async function fetchEvents() {
        try {
            const res = await request("/events", "GET", null, { Authorization: `Bearer ${cookies.token}` });
            // Ensure dates are Date objects
            events = res.map(e => ({ ...e, date: new Date(e.date) }));
        } catch (e) {
            console.error(e);
        }
    }

    function prevMonth() {
        currentDate = new Date(year, month - 1, 1);
    }

    function nextMonth() {
        currentDate = new Date(year, month + 1, 1);
    }

    function selectDate(day) {
        selectedDate = new Date(year, month, day);
        eventTitle = "";
        eventDesc = "";
        editingEventId = null;
        showModal = true;
    }

    function getEventsForDay(day) {
        return events.filter(e => {
            return e.date.getDate() === day &&
                   e.date.getMonth() === month &&
                   e.date.getFullYear() === year;
        });
    }

    async function saveEvent() {
        const payload = JSON.stringify({
            title: eventTitle,
            description: eventDesc,
            date: selectedDate.toISOString(),
        });
        const headers = {
            Authorization: `Bearer ${cookies.token}`,
            "Content-Type": "application/json",
        };

        try {
            if (editingEventId) {
                await request(`/events/${editingEventId}`, "PUT", payload, headers);
            } else {
                await request("/events", "POST", payload, headers);
            }
            closeModal();
            fetchEvents();
        } catch (e) {
            console.error(e);
        }
    }

    async function deleteEvent(id) {
        if (!confirm("Delete this event?")) return;
        try {
            await request(`/events/${id}`, "DELETE", null, { Authorization: `Bearer ${cookies.token}` });
            fetchEvents();
            // If we are editing this event, close modal or reset form
            if (editingEventId === id) {
                closeModal();
            }
        } catch (e) {
            console.error(e);
        }
    }

    function editEvent(event) {
        eventTitle = event.title;
        eventDesc = event.description;
        editingEventId = event.id;
        // selectedDate is already set
    }

    function closeModal() {
        showModal = false;
        selectedDate = null;
        editingEventId = null;
    }

    onMount(fetchEvents);
</script>

<div class="space-y-8">
    <div class="flex items-center justify-between">
        <h1 class="text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-primary to-secondary">Calendar</h1>
        <div class="flex items-center gap-4 glass p-2 rounded-lg">
            <button on:click={prevMonth} class="p-2 hover:bg-white/10 rounded-lg text-white transition">&lt;</button>
            <h2 class="text-xl font-semibold text-white w-40 text-center">{monthName} {year}</h2>
            <button on:click={nextMonth} class="p-2 hover:bg-white/10 rounded-lg text-white transition">&gt;</button>
        </div>
    </div>

    <div class="grid grid-cols-7 gap-4 text-center text-gray-400 font-semibold mb-2">
        <div>Sun</div><div>Mon</div><div>Tue</div><div>Wed</div><div>Thu</div><div>Fri</div><div>Sat</div>
    </div>

    <div class="grid grid-cols-7 gap-4">
        {#each Array(firstDayOfMonth) as _}
            <div class="h-32 rounded-xl"></div>
        {/each}
        {#each Array(daysInMonth) as _, i}
            {@const day = i + 1}
            {@const dayEvents = getEventsForDay(day)}
            {@const isToday = day === new Date().getDate() && month === new Date().getMonth() && year === new Date().getFullYear()}
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <div
                class="h-32 glass-card p-2 cursor-pointer transition relative overflow-hidden hover:border-primary/50 group flex flex-col"
                class:border-primary={isToday}
                class:bg-primary-10={isToday}
                on:click={() => selectDate(day)}
            >
                <span class="font-bold text-gray-300 group-hover:text-white transition" class:text-primary={isToday}>{day}</span>
                <div class="mt-1 space-y-1 overflow-y-auto scrollbar-hide w-full flex-grow">
                    {#each dayEvents as event}
                        <div class="text-xs bg-primary/80 text-black font-semibold px-2 py-1 rounded shadow-sm truncate w-full block" title={event.title}>
                            {event.title}
                        </div>
                    {/each}
                </div>
            </div>
        {/each}
    </div>
</div>

{#if showModal}
    <div class="fixed inset-0 bg-black/60 backdrop-blur-sm flex items-center justify-center z-50 p-4">
        <div class="glass-card p-8 rounded-2xl shadow-2xl w-full max-w-md border border-white/10">
            <h3 class="text-2xl font-bold text-white mb-6 border-b border-white/10 pb-2">
                {selectedDate?.toLocaleDateString(undefined, { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' })}
            </h3>

            <!-- List existing events for this day to edit/delete -->
            {#if getEventsForDay(selectedDate.getDate()).length > 0}
                <div class="mb-6 space-y-3">
                    <h4 class="text-xs font-bold text-gray-400 uppercase tracking-wider">Events</h4>
                    {#each getEventsForDay(selectedDate.getDate()) as event}
                        <div class="bg-white/5 p-3 rounded-lg flex justify-between items-center border border-white/5 hover:border-white/10 transition">
                            <div>
                                <div class="font-bold text-white">{event.title}</div>
                                <div class="text-xs text-gray-400">{event.description}</div>
                            </div>
                            <div class="flex gap-2">
                                <button on:click|stopPropagation={() => editEvent(event)} class="text-primary hover:text-green-300 text-sm font-medium">Edit</button>
                                <button on:click|stopPropagation={() => deleteEvent(event.id)} class="text-accent hover:text-red-300 text-sm font-medium">Del</button>
                            </div>
                        </div>
                    {/each}
                </div>
            {/if}

            <h4 class="text-xs font-bold text-gray-400 uppercase tracking-wider mb-3">{editingEventId ? "Edit Event" : "Add Event"}</h4>
            <div class="space-y-4">
                <input
                    bind:value={eventTitle}
                    placeholder="Event Title"
                    class="w-full glass-input px-4 py-3 rounded-lg focus:ring-2 focus:ring-primary transition"
                />
                <textarea
                    bind:value={eventDesc}
                    placeholder="Description"
                    rows="3"
                    class="w-full glass-input px-4 py-3 rounded-lg focus:ring-2 focus:ring-primary transition resize-none"
                ></textarea>
                <div class="flex justify-end gap-3 mt-6">
                    <button on:click={closeModal} class="px-4 py-2 text-gray-300 hover:text-white transition">Close</button>
                    <button on:click={saveEvent} class="glass-button px-6 py-2 rounded-lg hover:shadow-lg hover:shadow-primary/20 transition transform hover:-translate-y-0.5">Save</button>
                </div>
            </div>
        </div>
    </div>
{/if}
