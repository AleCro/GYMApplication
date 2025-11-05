<script>
  import { page } from '$app/stores';
  import favicon from '$lib/assets/favicon.svg';

  export let data;

  let navLinks = [
    { label: 'Home', href: '/' },
    { label: 'Notes', href: '/notes' },
    { label: 'Calendar', href: '/calendar' },
    { label: 'Exercise', href: '/exercise' },
    { label: 'Progress', href: '/progress' },
    { label: 'Goals', href: '/goals' },
  ];

  // reactive current path so Svelte updates on client-side navigation
  $: currentPath = $page.url?.pathname ?? '/';

  // normalize paths to avoid mismatch from trailing slashes
  const normalize = (p) => {
    if (!p) return '/';
    if (p === '/') return '/';
    return p.endsWith('/') ? p.slice(0, -1) : p;
  };

  const isActive = (href) => {
    const a = normalize(href);
    const b = normalize(currentPath);
    return a === b || (a !== '/' && b.startsWith(a));
  };
</script>

<link rel="stylesheet" href="/app.css">

<svelte:head>
  <link rel="icon" href={favicon} />
  <title>GYM App</title>
</svelte:head>

<nav class="navbar navbar-expand-lg navbar-dark bg-dark border-bottom border-secondary">
  <div class="container">
    <a class="navbar-brand d-flex align-items-center gap-2" href="/">
      <img src={favicon} alt="GYM" width="34" height="34" class="rounded" />
      <div class="d-flex flex-column">
        <strong class="m-0">Ale's GYM App</strong>
        <small class="text-muted" style="font-size: 10px">Workout efficiently <sup>(unlike this site)</sup></small>
      </div>
    </a>

    <button
      class="navbar-toggler"
      type="button"
      data-bs-toggle="collapse"
      data-bs-target="#mainNavbar"
      aria-controls="mainNavbar"
      aria-expanded="false"
      aria-label="Toggle navigation"
    >
      <span class="navbar-toggler-icon"></span>
    </button>

    <div class="collapse navbar-collapse" id="mainNavbar">
      <ul class="navbar-nav ms-3 me-auto mb-2 mb-lg-0">
        {#each navLinks as link}
          <li class="nav-item">
            <a
              class="nav-link"
              class:active={isActive(link.href)}
              href={link.href}
              sveltekit:prefetch
              aria-current={isActive(link.href) ? 'page' : undefined}
            >
              {link.label}
            </a>
          </li>
        {/each}
      </ul>

      <div class="d-flex gap-2">
        <a class="btn btn-sm btn-outline-light" href="/login">Sign in</a>
      </div>
    </div>
  </div>
</nav>

<main class="container my-4">
  <div class="dark-panel p-4 rounded">
    <slot />
  </div>
</main>

<footer class="footer-dark mt-4">
  <div class="container py-3 d-flex justify-content-between align-items-center">
    <div class="small text-muted">© {new Date().getFullYear()} Alejandro's GYM App</div>
    <div class="small text-muted"><a href="https://github.com/AleCro/GYMApplication">GYM App</a> · Built with ♥</div>
  </div>
</footer>

<style>
  :global(body) {
    background: linear-gradient(180deg,#06070a 0%, #0b0f14 100%);
    color: #e6eef6;
    -webkit-font-smoothing: antialiased;
    font-family: system-ui, -apple-system, "Segoe UI", Roboto, "Helvetica Neue", Arial;
  }

  .navbar {
    background: rgba(8,10,13,0.85) !important;
    backdrop-filter: blur(6px);
  }

  .navbar .navbar-brand img { object-fit: cover; filter: drop-shadow(0 2px 6px rgba(0,0,0,.6)); }

  .nav-link {
    color: rgba(230,238,246,0.8) !important;
    transition: color .15s ease, transform .12s ease;
  }

  .nav-link:hover { color: #fff !important; transform: translateY(-1px); }

  .nav-link.active {
    color: #9be7ff !important;
    font-weight: 600;
  }

  .nav-link.active::after {
    content: '';
    display: block;
    height: 3px;
    width: 40%;
    background: linear-gradient(90deg, rgba(155,231,255,0.95), rgba(124,58,237,0.9));
    border-radius: 3px;
    margin-top: 6px;
    opacity: 0.95;
  }

  .dark-panel {
    background: linear-gradient(180deg, rgba(20,24,30,0.9), rgba(12,14,18,0.85));
    border: 1px solid rgba(255,255,255,0.04);
    box-shadow: 0 8px 30px rgba(2,6,23,0.6);
    color: #e6eef6;
    min-height: 60vh;
  }

  .footer-dark {
    background: linear-gradient(0deg, rgba(6,8,12,0.6), rgba(8,10,14,0.8));
    border-top: 1px solid rgba(255,255,255,0.03);
    color: rgba(230,238,246,0.7);
  }

  footer a {
    color: rgba(155,231,255,0.8);
    text-decoration: none;
  }
  footer a:hover {
    text-decoration: underline;
    color: #9be7ff;
  }

  /* small responsive tweaks */
  @media (max-width: 576px) {
    .navbar-brand .d-flex { display: none; } /* hide brand text on tiny screens */
    .dark-panel { padding: 1rem; }
  }
</style>
