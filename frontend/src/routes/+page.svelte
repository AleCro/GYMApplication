<script>
  import { fly } from 'svelte/transition';

  // If you're using SvelteKit, page props land in `data` when provided by a loader.
  // This component also works if you pass `data` in manually.
  export let data = {};

  // safe helpers (copied / adapted from your app)
  const MS_THRESHOLD = 1e12;
  function toMs(ts) {
    const n = Number(ts);
    if (Number.isNaN(n)) return null;
    return n > MS_THRESHOLD ? n : n * 1000;
  }

  function formatTimeShort(d) {
    return d.toLocaleString([], { month: 'short', day: 'numeric', hour: 'numeric', minute: '2-digit' });
  }

  // user-safe getters
  const userName = data?.user?.username || 'Athlete';
  const notesArr = (data?.user?.notes && Array.isArray(data.user.notes)) ? data.user.notes : [];
  const exercisesArr = (data?.user?.exercises && Array.isArray(data.user.exercises)) ? data.user.exercises : [];

  // parse calendar events if available
  let events = [];
  if (data?.user?.calendar && Array.isArray(data.user.calendar)) {
    events = data.user.calendar
      .map((e, i) => {
        const ms = toMs(e.time);
        if (!ms) return null;
        return {
          // keep original fields, plus normalized time + fallback id
          id: e.id ?? `${ms}-${i}`,
          name: e.name ?? e.title ?? 'Event',
          timeMs: ms,
          date: new Date(ms)
        };
      })
      .filter(Boolean)
      .sort((a, b) => a.timeMs - b.timeMs);
  }

  const now = new Date();
  const upcoming = events.filter(ev => ev.date >= now).slice(0, 3);

  // quick stats
  const notesCount = notesArr.length;
  const eventsCount = events.length;
  const exercisesCount = exercisesArr.length;
</script>

<div class="welcome-container">
  <section class="hero row align-items-center">
    <div class="col-lg-7">
      <div in:fly={{ y: -8, duration: 300 }}>
        <h1 class="display-5 mb-2">Welcome back{userName ? `, ${userName}` : ''} 👋</h1>
        <p class="lead text-muted mb-3">Ready for today's session? Quick access to your notes, upcoming events, and exercises.</p>

        <div class="d-flex gap-2 flex-wrap">
          <a class="btn btn-primary btn-lg" href="/calendar">View Calendar</a>
          <a class="btn btn-outline-light btn-lg" href="/notes">Open Notes</a>
          <a class="btn btn-outline-light btn-lg" href="/exercise">Exercises</a>
        </div>
      </div>
    </div>
  </section>

  <section class="mt-4">
    <div class="row g-3">
      <div class="col-lg-8">
              <div class="card summary-card p-3 h-100" in:fly={{ y: 6, duration: 300 }}>
        <div class="d-flex justify-content-between align-items-start mb-2">
          <div>
            <div class="small text-muted">Quick Overview</div>
            <h5 class="mb-0">Today at a glance</h5>
          </div>
        </div>

        <div class="row g-2 mt-3">
          <div class="col-4">
            <div class="stat-box">
              <div class="stat-value">{notesCount}</div>
              <div class="stat-label">Notes</div>
            </div>
          </div>
          <div class="col-4">
            <div class="stat-box">
              <div class="stat-value">{eventsCount}</div>
              <div class="stat-label">Events</div>
            </div>
          </div>
          <div class="col-4">
            <div class="stat-box">
              <div class="stat-value">{exercisesCount}</div>
              <div class="stat-label">Exercises</div>
            </div>
          </div>
        </div>

        <hr class="my-3" />

        <div>
          <div class="small text-muted mb-2">Upcoming</div>

          {#if upcoming.length === 0}
            <div class="text-muted small">No upcoming events.</div>
          {:else}
            <ul class="list-unstyled mb-0">
              {#each upcoming as ev (ev.id)}
                <li class="upcoming-item">
                  <div class="fw-semibold">{ev.name}</div>
                  <div class="small text-muted">{formatTimeShort(ev.date)}</div>
                </li>
              {/each}
            </ul>
          {/if}
        </div>
      </div>
      </div>

      <div class="col-lg-4">
        <div class="card panels p-3">
          <h6 class="mb-3">Quick Actions</h6>
          <div class="d-grid gap-2">
            <a class="btn btn-outline-light" href="/notes">New Note</a>
            <a class="btn btn-outline-light" href="/calendar">Add Event</a>
            <a class="btn btn-outline-light" href="/exercise">Start Workout</a>
          </div>
        </div>
      </div>
    </div>
  </section>
</div>

<style>
  .welcome-container { color: #e6eef6; }

  .hero { gap: 24px; }
  h1.display-5 { font-weight: 700; color: #fff; letter-spacing: -0.4px; }
  p.lead { color: rgba(230,238,246,0.75); }

  .summary-card {
    background: linear-gradient(180deg, rgba(255,255,255,0.02), rgba(255,255,255,0.01));
    border: 1px solid rgba(255,255,255,0.04);
    box-shadow: 0 6px 28px rgba(2,6,23,0.6);
    min-height: 210px;
  }

  .stat-box {
    background: rgba(255,255,255,0.02);
    border-radius: .5rem;
    padding: 10px;
    text-align: center;
  }
  .stat-value {
    font-size: 1.1rem;
    font-weight: 700;
    color: #9be7ff;
  }
  .stat-label { font-size: 0.8rem; color: rgba(230,238,246,0.7); }

  .upcoming-item {
    padding: 8px 0;
    border-bottom: 1px dashed rgba(255,255,255,0.03);
  }
  .upcoming-item:last-child { border-bottom: none; }

  .panels {
    background: linear-gradient(180deg, rgba(20,24,30,0.9), rgba(12,14,18,0.85));
    border: 1px solid rgba(255,255,255,0.03);
  }

  .note-summary {
    padding: 10px 0;
    border-bottom: 1px dashed rgba(255,255,255,0.03);
  }
  .note-summary:last-child { border-bottom: none; }

  .link-light { color: #9be7ff; text-decoration: none; }
  .link-light:hover { text-decoration: underline; color: #bff1ff; }

  a.btn-primary { background: linear-gradient(90deg,#0d6efd,#7c3ef3); border: none; }
  a.btn-outline-light { border-color: rgba(255,255,255,0.06); color: rgba(230,238,246,0.95); }

  @media (max-width: 767px) {
    .stat-value { font-size: 1rem; }
    .summary-card { min-height: auto; }
  }
</style>
