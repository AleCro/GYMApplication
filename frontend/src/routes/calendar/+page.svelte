<script>
	import { onMount } from 'svelte';

	export let data = {
		user: {
			calendar: [
				{ name: 'Doctor Appointment', time: 1760481795952 },
				{ name: 'Team Meeting', time: 1760475600000 },
				{ name: 'Birthday Party', time: 1760562000000 }
			]
		}
	};

	const MS_THRESHOLD = 1e12;

	function toMs(ts) {
		const n = Number(ts);
		if (Number.isNaN(n)) return null;
		return n > MS_THRESHOLD ? n : n * 1000;
	}

	function sameLocalDate(a, b) {
		return (
			a.getFullYear() === b.getFullYear() &&
			a.getMonth() === b.getMonth() &&
			a.getDate() === b.getDate()
		);
	}

	function startOfToday() {
		const t = new Date();
		t.setHours(0, 0, 0, 0);
		return t;
	}

	let events = [];
	let displayedMonth;
	let displayedYear;
	let calendar = [];
	const today = new Date();

	function formatTime(date) {
		return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
	}

	function parseEvents() {
		const raw = data && data.user && Array.isArray(data.user.calendar) ? data.user.calendar : [];
		events = raw
			.map((e, i) => {
				const ms = toMs(e.time);
				if (!ms) return null;
				const d = new Date(ms);
				// create a stable unique id per input index (ms + index)
				const id = `${ms}-${i}`;
				return { ...e, timeMs: ms, date: d, id };
			})
			.filter(Boolean)
			.sort((a, b) => a.timeMs - b.timeMs);
	}

	function chooseInitialMonthYear() {
		if (events.length === 0) {
			displayedMonth = today.getMonth();
			displayedYear = today.getFullYear();
			return;
		}
		const now = startOfToday();
		const upcoming = events.find((ev) => ev.date >= now);
		const choose = upcoming || events[0];
		displayedMonth = choose.date.getMonth();
		displayedYear = choose.date.getFullYear();
	}

	function buildCalendar() {
		calendar = [];
		const firstOfMonth = new Date(displayedYear, displayedMonth, 1);
		const firstWeekday = firstOfMonth.getDay();
		const daysInMonth = new Date(displayedYear, displayedMonth + 1, 0).getDate();

		let dayCounter = 1;
		for (let week = 0; week < 6; week++) {
			const weekRow = [];
			for (let wd = 0; wd < 7; wd++) {
				if (week === 0 && wd < firstWeekday) {
					weekRow.push(null);
				} else if (dayCounter > daysInMonth) {
					weekRow.push(null);
				} else {
					weekRow.push(new Date(displayedYear, displayedMonth, dayCounter));
					dayCounter++;
				}
			}
			calendar.push(weekRow);
			if (dayCounter > daysInMonth && week >= 4) break;
		}
	}

	function getEventsForDate(dateObj) {
		return events.filter((ev) => sameLocalDate(ev.date, dateObj));
	}

	function prevMonth() {
		if (displayedMonth === 0) {
			displayedMonth = 11;
			displayedYear -= 1;
		} else {
			displayedMonth -= 1;
		}
		buildCalendar();
	}
	function nextMonth() {
		if (displayedMonth === 11) {
			displayedMonth = 0;
			displayedYear += 1;
		} else {
			displayedMonth += 1;
		}
		buildCalendar();
	}
	function goToToday() {
		displayedMonth = today.getMonth();
		displayedYear = today.getFullYear();
		buildCalendar();
	}

	// Only reinitialize when `data` actually changes (avoid resetting when user navigates months)
	let _lastDataJson = '';
	$: {
		const json = JSON.stringify(data || {});
		if (json !== _lastDataJson) {
			_lastDataJson = json;
			parseEvents();
			chooseInitialMonthYear();
			buildCalendar();
		}
	}

	onMount(() => {
		// initial parse in case data was set before mount (defensive)
		parseEvents();
		chooseInitialMonthYear();
		buildCalendar();
	});

	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	// Form state
	let title = '';
	// default date to today (YYYY-MM-DD)
	let dateValue = new Date().toISOString().slice(0, 10);
	// default time to current hour:minute
	const pad = (n) => String(n).padStart(2, '0');
	const now = new Date();
	let timeValue = `${pad(now.getHours())}:${pad(now.getMinutes())}`;

	let loading = false;
	let successMsg = '';
	let errorMsg = '';

	// Basic client-side validation
	function validate() {
		if (!title.trim()) {
			errorMsg = 'Please enter a title for the event.';
			return false;
		}
		if (!dateValue) {
			errorMsg = 'Please pick a date.';
			return false;
		}
		// time is optional; but we keep it required in UI
		return true;
	}

	// Convert local date + time string to millisecond timestamp
	function dateTimeToMs(dateStr, timeStr) {
		// If timeStr missing, set to 00:00
		const t = timeStr && timeStr.trim() ? timeStr : '00:00';
		// Create a local ISO-like string and use Date constructor
		// e.g. "2025-10-14T13:30"
		const iso = `${dateStr}T${t}`;
		const d = new Date(iso);
		return d.getTime();
	}

	async function submitForm(e) {
		e.preventDefault();
		successMsg = '';
		errorMsg = '';

		if (!validate()) return;

		const timeMs = dateTimeToMs(dateValue, timeValue);
		if (Number.isNaN(timeMs)) {
			errorMsg = 'Invalid date/time.';
			return;
		}

		const payload = {
			session: data.user.session,
			title: title.trim(),
			time: timeMs
		};

		loading = true;

		try {
			const res = await fetch('/calendar', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(payload)
			});

			if (!res.ok) {
				// try to parse a helpful message from the server response
				let text;
				try {
					text = await res.text();
				} catch (err) {
					text = res.statusText;
				}
				throw new Error(`Server responded ${res.status}: ${text}`);
			}

			let returned;
			try {
				returned = await res.json();
			} catch {
				returned = payload;
			}

			data.user.calendar.push({
				title: title.trim(),
				time: timeMs
			});

			data.user.calendar = data.user.calendar;

			buildCalendar();
			successMsg = 'Event saved.';
			dispatch('eventAdded', returned);
			title = '';
		} catch (err) {
			console.error(err);
			errorMsg = err.message || 'Failed to save event.';
		} finally {
			loading = false;
		}
	}
</script>

<div class="container mt-4">
	<div class="d-flex justify-content-between align-items-center mb-2">
		<div>
			<button class="btn btn-outline-primary me-1" on:click={prevMonth} aria-label="Previous month"
				>&laquo; Prev</button
			>
			<button class="btn btn-outline-primary me-2" on:click={nextMonth} aria-label="Next month"
				>Next &raquo;</button
			>
			<button class="btn btn-secondary" on:click={goToToday}>Today</button>
		</div>

		<h3 class="mb-0">
			{new Date(displayedYear, displayedMonth).toLocaleString('default', { month: 'long' })}
			{displayedYear}
		</h3>

		<div>
			<span class="small text-muted"
				>{events.length} event{events.length !== 1 ? 's' : ''} total</span
			>
		</div>
	</div>

	<div class="table-responsive">
		<table class="table table-bordered text-center">
			<thead class="table-light">
				<tr>
					<th>Sun</th><th>Mon</th><th>Tue</th><th>Wed</th>
					<th>Thu</th><th>Fri</th><th>Sat</th>
				</tr>
			</thead>
			<tbody>
				{#each calendar as week}
					<tr>
						{#each week as day}
							{#if !day}
								<td class="align-top" style="height:110px; vertical-align:top;"></td>
							{:else}
								<td class="align-top" style="height:110px; vertical-align:top; width:14%;">
									<div class="d-flex justify-content-between align-items-start">
										<small class="text-muted">{day.getDate()}</small>
										{#if sameLocalDate(day, today)}
											<span class="badge bg-success">Today</span>
										{/if}
									</div>

									<div class="mt-2" style="max-height:65px; overflow:auto;">
										{#each getEventsForDate(day) as ev (ev.id)}
											<div class="card mb-1" style="font-size:0.85rem;">
												<div class="card-body p-2">
													<div class="fw-semibold">{ev.title}</div>
													<div class="small text-muted">{formatTime(ev.date)}</div>
												</div>
											</div>
										{/each}
									</div>
								</td>
							{/if}
						{/each}
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
</div>

<form on:submit|preventDefault={submitForm} class="card p-3 mb-3 container mt-4">
	<div class="mb-3">
		<label for="title" class="form-label">Title</label>
		<input id="title" class="form-control" placeholder="Event title" bind:value={title} required />
	</div>

	<div class="row g-2 mb-3">
		<div class="col-md-6">
			<label for="date" class="form-label">Date</label>
			<input id="date" type="date" class="form-control" bind:value={dateValue} required />
		</div>

		<div class="col-md-6">
			<label for="time" class="form-label">Time</label>
			<input id="time" type="time" class="form-control" bind:value={timeValue} />
		</div>
	</div>

	<div class="d-flex gap-2">
		<button type="submit" class="btn btn-primary" disabled={loading}>
			{#if loading}
				<span class="spinner-border spinner-border-sm me-2" role="status" aria-hidden="true"></span>
				Saving...
			{:else}
				Add event
			{/if}
		</button>

		<button
			type="button"
			class="btn btn-secondary"
			on:click={() => {
				title = '';
			}}
		>
			Clear title
		</button>
	</div>

	{#if successMsg}
		<div class="alert alert-success mt-3" role="alert">{successMsg}</div>
	{/if}
	{#if errorMsg}
		<div class="alert alert-danger mt-3" role="alert">{errorMsg}</div>
	{/if}
</form>

<style>
	table.table td {
		padding: 6px;
	}
	.card {
		border-radius: 0.35rem;
	}
	form.card {
		max-width: 720px;
	}
</style>
