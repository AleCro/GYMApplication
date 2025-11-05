<script>
	import { onMount, createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();

	export let data = {
		user: {
			session: '',
			calendar: []
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

	function getToday() {
		const t = new Date();
		t.setHours(0, 0, 0, 0);
		return t;
	}

	let events = [];
	let displayedMonth;
	let displayedYear;
	let calendar = [];

	function formatTime(date, timezone) {
		return new Intl.DateTimeFormat('en-US', {
			hour: '2-digit',
			minute: '2-digit',
			timeZone: timezone || Intl.DateTimeFormat().resolvedOptions().timeZone
		}).format(date);
	}

	function parseEvents() {
		const raw = data && data.user && Array.isArray(data.user.calendar)
			? data.user.calendar
			: [];

		events = raw
			.map((e, i) => {
				const ms = toMs(e.time);
				if (!ms) return null;
				const d = new Date(ms);
				const id = `${ms}-${i}`;
				return {
					...e,
					title: e.title || e.name,
					timeMs: ms,
					date: d,
					id,
					originalIndex: i
				};
			})
			.filter(Boolean)
			.sort((a, b) => a.timeMs - b.timeMs);
	}

	function chooseInitialMonthYear() {
		if (events.length === 0) {
			const today = getToday();
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
		const t = getToday();
		displayedMonth = t.getMonth();
		displayedYear = t.getFullYear();
		buildCalendar();
	}

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
		parseEvents();
		chooseInitialMonthYear();
		buildCalendar();
	});

	let title = '';
	const pad = (n) => String(n).padStart(2, '0');
	const now = new Date();
	let dateValue = `${now.getFullYear()}-${pad(now.getMonth() + 1)}-${pad(now.getDate())}`;
	let timeValue = `${pad(now.getHours())}:${pad(now.getMinutes())}`;
	let loading = false;
	let successMsg = '';
	let errorMsg = '';

	function validate() {
		if (!title.trim()) {
			errorMsg = 'Please enter a title for the event.';
			return false;
		}
		if (!dateValue) {
			errorMsg = 'Please pick a date.';
			return false;
		}
		return true;
	}

	function dateTimeToMs(dateStr, timeStr) {
		const [year, month, day] = dateStr.split('-').map(Number);
		const [hour, minute] = (timeStr || '00:00').split(':').map(Number);
		const localDate = new Date(year, month - 1, day, hour, minute, 0, 0);
		const utcMs = localDate.getTime() - localDate.getTimezoneOffset() * 60000;
		return utcMs;
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

		//  Include timezone in payload
		const payload = {
			session: data.user.session,
			title: title.trim(),
			time: timeMs,
			timezone: Intl.DateTimeFormat().resolvedOptions().timeZone
		};

		loading = true;

		try {
			const res = await fetch('/calendar', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify(payload)
			});

			if (!res.ok) throw new Error('Failed to save event.');

			data.user.calendar.push({
				title: title.trim(),
				time: timeMs,
				timezone: payload.timezone
			});

			data.user.calendar = data.user.calendar;
			buildCalendar();
			successMsg = 'Event saved.';
			dispatch('eventAdded', payload);
			title = '';
		} catch (err) {
			errorMsg = err.message || 'Failed to save event.';
		} finally {
			loading = false;
		}
	}

	async function deleteEvent(index) {
		if (!confirm('Delete this event?')) return;

		try {
			const res = await fetch('/deleteevent', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ session: data.user.session, i: index })
			});

			if (!res.ok) throw new Error('Failed to delete event.');

			data.user.calendar.splice(index, 1);
			data.user.calendar = data.user.calendar;
			buildCalendar();
		} catch (err) {
			alert(err || 'Error deleting event.');
		}
	}

	function setToday() {
		const t = getToday();
		const pad = (n) => String(n).padStart(2, '0');
		dateValue = `${t.getFullYear()}-${pad(t.getMonth() + 1)}-${pad(t.getDate())}`;
		timeValue = `${pad(new Date().getHours())}:${pad(new Date().getMinutes())}`;
		displayedMonth = t.getMonth();
		displayedYear = t.getFullYear();
		buildCalendar();
	}
</script>

<div class="container mt-4">
	<div class="d-flex justify-content-between align-items-center mb-2">
		<div>
			<button class="btn btn-outline-primary me-1" on:click={prevMonth}>&laquo; Prev</button>
			<button class="btn btn-outline-primary me-2" on:click={nextMonth}>Next &raquo;</button>
			<button class="btn btn-secondary" on:click={goToToday}>Today</button>
		</div>

		<h3 class="mb-0">
			{new Date(displayedYear, displayedMonth).toLocaleString('default', { month: 'long' })} {displayedYear}
		</h3>

		<div>
			<span class="small text-muted">{events.length} event{events.length !== 1 ? 's' : ''} total</span>
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
								<td class="align-top" style="height:110px;"></td>
							{:else}
								<td class="align-top" style="height:110px; vertical-align:top; width:14%;">
									<div class="d-flex justify-content-between align-items-start">
										<small class="text-muted">{day.getDate()}</small>
										{#if sameLocalDate(day, getToday())}
											<span class="badge bg-success">Today</span>
										{/if}
									</div>

									<div class="mt-2" style="max-height:65px; overflow:auto;">
										{#each getEventsForDate(day) as ev (ev.id)}
											<div class="card mb-1" style="font-size:0.85rem;">
												<div class="card-body p-2 d-flex justify-content-between align-items-center">
													<div>
														<div class="fw-semibold">{ev.title}</div>
														<div class="small text-muted">
															{formatTime(ev.date, ev.timezone)}
														</div>
													</div>
													<button
														class="btn btn-sm btn-outline-danger ms-2"
														title="Delete event"
														on:click={() => deleteEvent(ev.originalIndex)}
													>
														🗑️
													</button>
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

	<div class="d-flex justify-content-between align-items-center mt-3">
		<button type="button" class="btn btn-outline-primary" on:click={prevMonth}>
			&laquo; Prev
		</button>

		<div class="d-flex justify-content-center gap-3">
			<button type="submit" class="btn btn-primary" disabled={loading}>
				{#if loading}
					<span class="spinner-border spinner-border-sm me-2" role="status"></span> Saving...
				{:else}
					Add Event
				{/if}
			</button>

			<button type="button" class="btn btn-secondary" on:click={() => (title = '')}>
				Clear Title
			</button>

			<button type="button" class="btn btn-outline-info" on:click={setToday}>
				Today
			</button>
		</div>

		<button type="button" class="btn btn-outline-primary" on:click={nextMonth}>
			Next &raquo;
		</button>
	</div>

	{#if successMsg}
		<div class="alert alert-success mt-3">{successMsg}</div>
	{/if}
	{#if errorMsg}
		<div class="alert alert-danger mt-3">{errorMsg}</div>
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
