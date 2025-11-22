<script>
  export const layout = false;
</script>

<svelte:head>
  <title>About</title>
</svelte:head>

<div class="min-h-screen flex bg-gray-900 text-gray-100">
  <aside
    class="hidden lg:flex flex-col w-72 bg-gray-800 border-r border-gray-700 p-8 sticky top-0 h-screen"
  >
    <div class="space-y-6 flex-1">
      <a
        href="/"
        class="flex items-center gap-3 focus:outline-none focus-visible:ring-2 focus-visible:ring-indigo-500 rounded-lg -m-1 p-1"
      >
        <div
          class="w-10 h-10 rounded-lg bg-gradient-to-br from-indigo-500 to-pink-500 flex items-center justify-center text-white font-bold text-xl shadow-md shrink-0"
        >
          AG
        </div>
        <h2 class="text-2xl font-bold">About</h2>
      </a>
      <br />

      <section
        class="p-4 rounded-lg bg-gray-700 border border-gray-600 text-sm leading-relaxed"
      >
        <h3 class="font-semibold text-white mb-2">About</h3>
        <p class="text-gray-300">
          About AleGYM
        </p>
      </section>
    </div>

    <footer
      class="pt-8 border-t border-gray-700 text-xs text-gray-500 space-y-1 mt-auto"
    >
      <button
        on:click={() => history.back()}
        class="flex items-center gap-2 mt-4 px-3 py-2 -ml-3 rounded-lg text-sm font-medium text-indigo-400 hover:bg-gray-700 hover:text-indigo-300 transition focus-visible:ring-2 focus-visible:ring-indigo-500"
      >
        <svg
          class="w-5 h-5"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <polyline
            points="15 18 9 12 15 6"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
        Return to Application
      </button>
    </footer>
  </aside>

  <main class="flex-grow p-8 lg:p-12 overflow-y-auto">
    <article class="space-y-10 max-w-4xl mx-auto leading-relaxed">
      <header class="pb-4 border-b border-gray-700">
        <h1 class="text-4xl font-extrabold text-white">About</h1>
      </header>

      <section>
        <p>
          GYMApp is a fitness companion that unifies everyday training tasks
          into a clean, intuitive experience. A Svelte front end paired with a
          Go REST API and MongoDB supports a dashboard where users jot quick
          notes during workouts, schedule routines on an interactive calendar,
          and log progress with weight, messages, and optional photos stored in
          Base64.
        </p>

        <p>
          An exercise explorer offers examples and guidance to help users verify
          form and discover new workouts. The Goals and Progress features allow
          you to stay organized and on track throughout your fitness journey.
        </p>

        <p>
          By combining planning, reflection, and guidance in one place, GYMApp
          helps you stay focused on your training routine and long-term fitness
          goals.
        </p>
      </section>
    </article>
  </main>
</div>

<style>
  .about-container {
    max-width: 800px;
    margin: 2rem auto;
    padding: 1rem 1.5rem;
    line-height: 1.6;
    font-size: 1.05rem;
  }

  h1 {
    margin-bottom: 1rem;
    font-size: 2rem;
    font-weight: 700;
  }

  p {
    margin-bottom: 1rem;
  }
</style>
