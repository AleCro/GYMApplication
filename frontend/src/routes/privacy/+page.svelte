<script>
    // 1. SvelteKit directive to disable parent layout
    export const layout = false;

    // 2. Company-specific data
    const company = {
        name: 'AleGYM',
        contact: 'privacy@alegym.com', // Updated example email
        address: '123 Health & Fitness Blvd, Svelte City, GO 10101', // Updated example address
        effectiveDate: 'November 17, 2025' // Using today's date
    };

    // Helper for smooth scrolling to sections
    function scrollToSection(id) {
        document.getElementById(id)?.scrollIntoView({ behavior: 'smooth' });
    }
</script>

<svelte:head>
    <title>Privacy Policy - {company.name}</title>
    <meta name="description" content="The official Privacy Policy for {company.name}, outlining how we collect, use, and protect your fitness, progress, and account data." />
    <meta property="og:title" content="Privacy Policy | {company.name}" />
    <meta property="og:description" content="Learn how {company.name} protects your personal and fitness-related data." />
    <meta property="og:type" content="website" />
</svelte:head>

<div class="min-h-screen flex bg-gray-900 text-gray-100">

    <aside class="hidden lg:flex flex-col w-72 bg-gray-800 border-r border-gray-700 p-8 sticky top-0 h-screen">
        <div class="space-y-6 flex-1">
            
            <a href="/" class="flex items-center gap-3 focus:outline-none focus-visible:ring-2 focus-visible:ring-indigo-500 rounded-lg -m-1 p-1">
                <div class="w-10 h-10 rounded-lg bg-gradient-to-br from-indigo-500 to-pink-500 flex items-center justify-center text-white font-bold text-xl shadow-md shrink-0">AG</div>
                <h2 class="text-2xl font-bold">Privacy Policy</h2>
            </a>
            <br>
            
            <section class="p-4 rounded-lg bg-gray-700 border border-gray-600 text-sm leading-relaxed">
                <h3 class="font-semibold text-white mb-2">Policy Overview</h3>
                <p class="text-gray-300">
                    This document details the collection, usage, and protection of your <strong>Account Data</strong> and <strong>Fitness Progress Data</strong> specific to the {company.name} service.
                </p>
            </section>

            <nav class="pt-4 border-t border-gray-700 space-y-2">
                <h4 class="text-sm font-semibold text-gray-400 uppercase tracking-wider mb-2">Sections</h4>
                <button on:click={() => scrollToSection('data-collected')} class="w-full text-left px-3 py-1 rounded-lg text-sm font-medium text-gray-300 hover:bg-gray-700 hover:text-white transition">1. Data We Collect</button>
                <button on:click={() => scrollToSection('data-use')} class="w-full text-left px-3 py-1 rounded-lg text-sm font-medium text-gray-300 hover:bg-gray-700 hover:text-white transition">2. How We Use Data</button>
                <button on:click={() => scrollToSection('data-sharing')} class="w-full text-left px-3 py-1 rounded-lg text-sm font-medium text-gray-300 hover:bg-gray-700 hover:text-white transition">3. Security & Retention</button>
                <button on:click={() => scrollToSection('your-rights')} class="w-full text-left px-3 py-1 rounded-lg text-sm font-medium text-gray-300 hover:bg-gray-700 hover:text-white transition">4. Your Privacy Rights</button>
                <button on:click={() => scrollToSection('contact')} class="w-full text-left px-3 py-1 rounded-lg text-sm font-medium text-gray-300 hover:bg-gray-700 hover:text-white transition">5. Contact Information</button>
            </nav>
        </div>

        <footer class="pt-8 border-t border-gray-700 text-xs text-gray-500 space-y-1 mt-auto">
            <p>Effective: {company.effectiveDate}</p>
            <button 
                on:click={() => history.back()}
                class="flex items-center gap-2 mt-4 px-3 py-2 -ml-3 rounded-lg text-sm font-medium text-indigo-400 hover:bg-gray-700 hover:text-indigo-300 transition focus-visible:ring-2 focus-visible:ring-indigo-500"
            >
                <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <polyline points="15 18 9 12 15 6" stroke-linecap="round" stroke-linejoin="round" />
                </svg>
                Return to Application
            </button>
        </footer>
    </aside>

    <main class="flex-grow p-8 lg:p-12 overflow-y-auto">
        <article class="space-y-10 max-w-4xl mx-auto leading-relaxed">
            
            <header class="pb-4 border-b border-gray-700">
                <h1 class="text-4xl font-extrabold text-white">Privacy Policy</h1>
                <p class="mt-2 text-sm text-gray-400">Effective Date: <strong>{company.effectiveDate}</strong></p>
            </header>

            <section class="space-y-4">
                <p class="lg:hidden text-center p-3 rounded-lg bg-gray-800 border border-gray-700 text-sm text-indigo-400">
                    <button on:click={() => history.back()} class="underline">&larr; Return to Application</button>
                </p>
                <p>
                    Welcome to <strong>{company.name}</strong> (“we,” “our,” or “us”), your dedicated fitness companion. Your privacy and the security of your health and fitness data are our top priority.
                    This Privacy Policy explains how we collect, use, and protect your <strong>personal and fitness-related information</strong> when you use the AleGYM application.
                </p>
            </section>

            <hr class="border-gray-700">

            <section id="data-collected">
                <h2 class="text-2xl font-bold text-indigo-400 mb-4">1. Information We Collect</h2>
                
                <h3 class="text-xl font-semibold text-gray-200 mt-6">A. Account and Security Data (User)</h3>
                <p>This information is collected when you create and log into your account:</p>
                <ul class="list-disc list-inside space-y-2 pl-4 text-gray-300">
                    <li><strong>Username:</strong> Your chosen identifier for the application.</li>
                    <li><strong>Password:</strong> Stored as a securely <strong>hashed</strong> value. We <strong>never</strong> store or have access to your plain-text password.</li>
                    <li><strong>User Group:</strong> A field to manage your account type (e.g., standard <strong>GroupUser</strong> or administrator <strong>GroupAdmin</strong>).</li>
                    <li><strong>Session Data:</strong> Temporary records used for secure authentication and login management.</li>
                </ul>

                <h3 class="text-xl font-semibold text-gray-200 mt-6">B. Fitness and Progress Data (User-Generated)</h3>
                <p>As a fitness tracking application, we store data you actively enter to track your health and progress. This data is <strong>highly personal and treated as sensitive information</strong>:</p>
                <ul class="list-disc list-inside space-y-2 pl-4 text-gray-300">
                    <li><strong>Notes:</strong> Your private text entries (Title, Content, creation/update times) for tracking thoughts, diet, or routines. (Corresponds to the <code>Note</code> structure).</li>
                    <li><strong>Calendar Events:</strong> Records of your scheduled activities (Title, Description, When) for planning workouts and routines. (Corresponds to the <code>CalendarEvent</code> structure).</li>
                    <li><strong>Progress Entries:</strong> Detailed records of your progress (Title, Description, creation/update times). Critically, this includes:
                        <ul class="list-[circle] list-inside ml-4 mt-1 text-gray-400">
                            <li><strong>ImageBase64:</strong> <strong>Base64-encoded image blobs</strong> (e.g., progress photos). This highly personal data is stored directly in the database. (Corresponds to the <code>ProgressEntry</code> structure).</li>
                        </ul>
                    </li>
                </ul>
            </section>

            <hr class="border-gray-700">

            <section id="data-use">
                <h2 class="text-2xl font-bold text-indigo-400 mb-4">2. How We Use Your Data</h2>
                <p>We use your data solely to provide, maintain, and improve the <strong>{company.name}</strong> service, specifically for:</p>
                <ul class="list-disc list-inside space-y-2 pl-4 text-gray-300">
                    <li><strong>Core Service Provision:</strong> To display your personal <strong>Notes</strong>, manage your <strong>Calendar Events</strong>, and show your <strong>Progress Entries</strong> (including your images) only to you.</li>
                    <li><strong>Account Authentication:</strong> To verify your identity and manage secure login sessions.</li>
                    <li><strong>Personalization:</strong> To tailor the app experience based on your history and goals.</li>
                    <li><strong>Service Maintenance:</strong> To debug issues, protect against errors, and ensure application stability.</li>
                </ul>
                <p class="mt-4 p-3 rounded-lg bg-red-900/20 text-red-300 border border-red-700/50">
                    <strong>Crucially:</strong> Your <strong>Notes</strong>, <strong>Calendar Events</strong>, and <strong>Progress Photos (ImageBase64)</strong> are <strong>private</strong> and are <strong>never shared publicly, with other users, or with third-party advertisers.</strong>
                </p>
            </section>

            <hr class="border-gray-700">

            <section id="data-sharing">
                <h2 class="text-2xl font-bold text-indigo-400 mb-4">3. Data Security and Retention</h2>

                <h3 class="text-xl font-semibold text-gray-200 mt-6">Data Persistence and Storage</h3>
                <p>
                    Your data, including your <strong>Progress Photos (ImageBase64)</strong>, is stored persistently in a database.
                    All personal records are linked to your unique <strong>User ID</strong> and managed internally within the AleGYM system.
                </p>

                <h3 class="text-xl font-semibold text-gray-200 mt-6">Data Sharing</h3>
                <p>
                    We <strong>do not sell your personal data</strong>, including your fitness records or images. We may share data with trusted <strong>third-party providers</strong> (e.g., cloud hosting) only as necessary to operate and maintain the application. These providers are bound by strict confidentiality agreements.
                </p>

                <h3 class="text-xl font-semibold text-gray-200 mt-6">Security Measures</h3>
                <p>
                    We employ robust technical and organizational measures to protect all data, especially sensitive fitness information:
                </p>
                <ul class="list-disc list-inside space-y-2 pl-4 text-gray-300">
                    <li><strong>Encryption:</strong> Data is encrypted both at rest and in transit (via SSL/TLS).</li>
                    <li><strong>Access Controls:</strong> Strict access policies are enforced to ensure that only you can retrieve or modify your personal fitness data.</li>
                    <li><strong>Password Hashing:</strong> Passwords are cryptographically <strong>hashed and salted</strong> using industry best practices.</li>
                </ul>
                
                <h3 class="text-xl font-semibold text-gray-200 mt-6">Data Retention</h3>
                <p>
                    We retain your account data and all associated fitness data (<strong>Notes, Calendar Events, Progress Entries, and Images</strong>) for as long as your account is active. Upon account deletion, all associated data is permanently removed from our systems, except where retention is legally required.
                </p>
            </section>

            <hr class="border-gray-700">

            <section id="your-rights">
                <h2 class="text-2xl font-bold text-indigo-400 mb-4">4. Your Privacy Rights</h2>
                <p>
                    Depending on your jurisdiction, you have the following rights concerning your data, which you can exercise at any time:
                </p>
                <ul class="list-disc list-inside space-y-2 pl-4 text-gray-300">
                    <li><strong>Right to Access:</strong> Request a copy of the personal data we hold about you.</li>
                    <li><strong>Right to Rectification:</strong> Request corrections to any incomplete or inaccurate data, such as updating a <strong>Progress Entry</strong> or <strong>Note</strong>.</li>
                    <li><strong>Right to Erasure (Right to be Forgotten):</strong> Request that we delete all your personal data, including your <strong>Progress Photos</strong>. This will result in account termination.</li>
                    <li><strong>Right to Data Portability:</strong> Request that your data be transferred to you or another service in a common, machine-readable format.</li>
                </ul>
                <p class="mt-4">To exercise these rights, please contact us using the details below.</p>
            </section>

            <hr class="border-gray-700">

            <section id="contact" class="pt-4">
                <h2 class="text-2xl font-bold text-indigo-400 mb-4">5. Contact Information</h2>
                <p>If you have any questions or concerns regarding this policy, please contact us:</p>
                <ul class="space-y-2 text-gray-300">
                    <li><strong>Email:</strong> <a href="mailto:{company.contact}" class="text-indigo-400 hover:text-indigo-300">{company.contact}</a></li>
                    <li><strong>Address:</strong> {company.address}</li>
                </ul>
            </section>

            <section class="pt-4 border-t border-gray-700 mt-10">
                <h2 class="text-2xl font-bold text-indigo-400 mb-4">6. Changes to This Policy</h2>
                <p>We reserve the right to modify this Privacy Policy at any time. We will notify you of significant changes by updating the “Effective Date” above or through other means. Continued use of the service after any updates constitutes your acceptance of the revised policy.</p>
            </section>
        </article>
    </main>
</div>