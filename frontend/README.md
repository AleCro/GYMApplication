# AleGYM Frontend Client

The AleGYM Frontend is a modern, responsive web application built to provide an engaging user experience for tracking gym activities. It leverages the power of SvelteKit for server-side rendering and static site generation, combined with TailwindCSS for rapid UI development.

## Technology Stack

*   **Framework:** [SvelteKit](https://kit.svelte.dev/) - The web framework for building Svelte apps.
*   **Language:** TypeScript / JavaScript
*   **Styling:** [TailwindCSS](https://tailwindcss.com/) - A utility-first CSS framework.
*   **Runtime/Package Manager:** [Bun](https://bun.sh/) (Recommended) or Node.js.
*   **Build Tool:** Vite.

## Features

*   **Responsive Design:** Optimized for both desktop and mobile devices.
*   **Interactive Dashboard:** Visualize your progress and upcoming events.
*   **Goal Tracking UI:** Intuitive interface for managing main goals and sub-goals.
*   **Secure Authentication:** Seamless login and session management integration with the backend API.

## Installation & Running

### Prerequisites
- Bun (preferred) or Node.js installed.
- The AleGYM API running locally or accessible via network.

### Steps

1.  **Navigate to the frontend directory:**
    ```bash
    cd frontend
    ```

2.  **Install Dependencies:**
    ```bash
    bun install
    # or
    npm install
    ```

3.  **Run Development Server:**
    ```bash
    bun run dev
    # or
    npm run dev
    ```
    The application will be available at `http://localhost:5173`.

## Building for Production

To create a production-ready build of the application:

```bash
bun run build
# or
npm run build
```

You can preview the production build locally using:

```bash
bun run preview
```

## Project Structure

*   `src/routes/`: Contains the file-system based routing for the application.
    *   `app/`: Main application routes (protected).
    *   `auth/`: Authentication routes (login, signup).
*   `src/lib/`: Shared components, utility functions, and stores.
*   `static/`: Static assets like images and fonts.

## License

This project is licensed under the MIT License.

Powered by [yxl-prz/YSvelGoK](https://github.com/yxl-prz/YSvelGoK).
