<script>
    import { onDestroy, setContext } from "svelte";

    // Optional: used by your layout header
    setContext("title", "Timer");

    // ----------------------------
    // MODE
    // ----------------------------
    let mode = "stopwatch"; // "stopwatch" | "countdown"

    // ============================
    // STOPWATCH (COUNT UP)
    // ============================
    // Elapsed time in milliseconds
    let swMs = 0;
    let swRunning = false;
    let swTimer = null;
    let swStartTime = 0; // for Date.now()-based timing

    function startStopwatch() {
        if (swRunning) return;
        swRunning = true;

        // When resuming, keep previous elapsed time
        swStartTime = Date.now() - swMs;

        // Update at 50ms instead of 20ms to reduce jitter
        swTimer = setInterval(() => {
            // Accurate elapsed time based on current time
            swMs = Date.now() - swStartTime;
        }, 50);
    }

    function pauseStopwatch() {
        swRunning = false;
        clearInterval(swTimer);
    }

    function resetStopwatch() {
        pauseStopwatch();
        swMs = 0;
    }

    function pad2(n) {
        return String(n).padStart(2, "0");
    }

    function pad3(n) {
        return String(n).padStart(3, "0");
    }

    // STOPWATCH FORMAT: MM:SS:MLL  (no hours)
    function formatSW() {
        const totalMs = swMs;
        const totalSeconds = Math.floor(totalMs / 1000);

        const mins  = Math.floor(totalSeconds / 60);
        const secs  = totalSeconds % 60;

        // Snap ms to nearest 10ms so last digit doesn't buzz like crazy
        const rawMs = totalMs % 1000; // 0–999
        const snappedMs = Math.floor(rawMs / 10) * 10; // 0,10,20,...,990

        return `${pad2(mins)}:${pad2(secs)}:${pad3(snappedMs)}`;
    }

    // ============================
    // COUNTDOWN (HH:MM:SS)
    // ============================

    // Inputs as strings so binding is clean
    let cdHours = "0";
    let cdMinutes = "1";
    let cdSeconds = "0";

    // Internal total seconds remaining
    let cdTotal = 0;
    let cdRunning = false;
    let cdTimer = null;

    function clampInt(value, min = 0, max = null) {
        let v = parseInt(value);
        if (isNaN(v)) v = 0;
        if (v < min) v = min;
        if (max !== null && v > max) v = max;
        return v;
    }

    function computeTotalFromInputs() {
        const h = clampInt(cdHours, 0);
        const m = clampInt(cdMinutes, 0, 59);
        const s = clampInt(cdSeconds, 0, 59);
        return h * 3600 + m * 60 + s;
    }

    function syncDisplayFromInputs() {
        if (!cdRunning) {
            cdTotal = computeTotalFromInputs();
        }
    }

    // Run once so initial display shows 00:01:00
    cdTotal = computeTotalFromInputs();

    function handleHoursInput(e) {
        cdHours = String(clampInt(e.target.value, 0));
        syncDisplayFromInputs();
    }

    function handleMinutesInput(e) {
        cdMinutes = String(clampInt(e.target.value, 0, 59));
        syncDisplayFromInputs();
    }

    function handleSecondsInput(e) {
        cdSeconds = String(clampInt(e.target.value, 0, 59));
        syncDisplayFromInputs();
    }

    function startCountdown() {
        if (cdRunning) return;

        cdTotal = computeTotalFromInputs();
        if (cdTotal <= 0) return; // nothing to count

        cdRunning = true;

        cdTimer = setInterval(() => {
            if (cdTotal > 0) {
                cdTotal--; // reactive → display updates
            } else {
                clearInterval(cdTimer);
                cdRunning = false;

                const audio = new Audio(
                    "https://actions.google.com/sounds/v1/alarms/beep_short.ogg"
                );
                audio.play().catch((e) => console.log("Audio playback failed:", e));
            }
        }, 1000);
    }

    function pauseCountdown() {
        cdRunning = false;
        clearInterval(cdTimer);
    }

    function resetCountdown() {
        pauseCountdown();
        cdTotal = computeTotalFromInputs();
    }

    function formatCD() {
        const total  = cdTotal;
        const hours  = Math.floor(total / 3600);
        const mins   = Math.floor((total % 3600) / 60);
        const secs   = total % 60;

        const hStr = pad2(hours);
        const mStr = pad2(mins);
        const sStr = pad2(secs);
        return `${hStr}:${mStr}:${sStr}`;
    }

    // Cleanup on destroy
    onDestroy(() => {
        clearInterval(swTimer);
        clearInterval(cdTimer);
    });
</script>

<!-- ============================
      TIMER UI
============================ -->
<div class="timer-wrapper max-w-xl mx-auto py-4">
    <!-- MODE SWITCH -->
    <div class="flex justify-center gap-4 mb-8">
        <button
            class="mode-btn {mode === 'stopwatch' ? 'active' : ''}"
            on:click={() => {
                mode = "stopwatch";
                resetCountdown();
            }}
        >
            Stopwatch
        </button>
        <button
            class="mode-btn {mode === 'countdown' ? 'active' : ''}"
            on:click={() => {
                mode = "countdown";
                resetStopwatch();
            }}
        >
            Countdown
        </button>
    </div>

    <div class="bg-white/5 rounded-2xl shadow-2xl p-8 border border-white/10">
        <!-- ============================
              STOPWATCH
        ============================ -->
        {#if mode === "stopwatch"}
            <div class="text-center">
                <h1 class="time-display text-primary">
                    {formatSW()}
                </h1>

                <div class="flex justify-center gap-4 mt-8">
                    <button class="btn green" on:click={startStopwatch} disabled={swRunning}>
                        ▶️ Start
                    </button>
                    <button class="btn yellow" on:click={pauseStopwatch} disabled={!swRunning}>
                        ⏸️ Pause
                    </button>
                    <button class="btn red" on:click={resetStopwatch} disabled={swMs === 0}>
                        🔄 Reset
                    </button>
                </div>
            </div>
        {/if}

        <!-- ============================
              COUNTDOWN
        ============================ -->
        {#if mode === "countdown"}
            <div class="text-center">
                <h1 class="time-display text-accent">
                    {formatCD()}
                </h1>

                <!-- HH : MM : SS inputs -->
                <div class="mt-6 flex items-center justify-center gap-2">
                    <div class="flex flex-col items-center">
                        <label class="text-gray-300 text-xs mb-1">Hours</label>
                        <input
                            type="number"
                            min="0"
                            bind:value={cdHours}
                            on:input={handleHoursInput}
                            class="input-field small"
                            disabled={cdRunning}
                        />
                    </div>

                    <span class="text-xl text-gray-300">:</span>

                    <div class="flex flex-col items-center">
                        <label class="text-gray-300 text-xs mb-1">Minutes</label>
                        <input
                            type="number"
                            min="0"
                            max="59"
                            bind:value={cdMinutes}
                            on:input={handleMinutesInput}
                            class="input-field small"
                            disabled={cdRunning}
                        />
                    </div>

                    <span class="text-xl text-gray-300">:</span>

                    <div class="flex flex-col items-center">
                        <label class="text-gray-300 text-xs mb-1">Seconds</label>
                        <input
                            type="number"
                            min="0"
                            max="59"
                            bind:value={cdSeconds}
                            on:input={handleSecondsInput}
                            class="input-field small"
                            disabled={cdRunning}
                        />
                    </div>
                </div>

                <div class="flex justify-center gap-4 mt-8">
                    <button
                        class="btn green"
                        on:click={startCountdown}
                        disabled={cdRunning || cdTotal <= 0}
                    >
                        ▶️ Start
                    </button>
                    <button class="btn yellow" on:click={pauseCountdown} disabled={!cdRunning}>
                        ⏸️ Pause
                    </button>
                    <button
                        class="btn red"
                        on:click={resetCountdown}
                        disabled={cdTotal === 0 && !cdRunning}
                    >
                        🔄 Reset
                    </button>
                </div>
            </div>
        {/if}
    </div>
</div>

<style>
    .time-display {
        font-size: clamp(3rem, 10vw, 6rem);
        font-weight: 800;
        letter-spacing: 2px;
        line-height: 1;
        font-variant-numeric: tabular-nums;
        /* Monospace-style digits to avoid shifting */
        font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono",
            "Courier New", monospace;
    }

    .mode-btn {
        padding: 0.6rem 1.5rem;
        font-size: 1rem;
        border-radius: 9999px;
        border: 1px solid rgba(255, 255, 255, 0.2);
        background: rgba(255, 255, 255, 0.08);
        color: white;
        transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    }
    .mode-btn:hover:not(.active) {
        background: rgba(255, 255, 255, 0.15);
    }
    .mode-btn.active {
        background: var(--primary-color, #6366f1);
        border-color: var(--primary-color, #6366f1);
        box-shadow: 0 0 15px rgba(99, 102, 241, 0.5);
        color: white;
    }

    .btn {
        padding: 0.8rem 1.5rem;
        font-size: 1.1rem;
        font-weight: 600;
        border-radius: 12px;
        color: white;
        transition: all 0.3s ease;
        display: flex;
        align-items: center;
        gap: 8px;
        box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
    }
    .btn:disabled {
        opacity: 0.5;
        cursor: not-allowed;
    }

    .btn.green { background: #22c55e; }
    .btn.green:hover:not(:disabled) { background: #16a34a; transform: translateY(-2px); }

    .btn.yellow { background: #eab308; }
    .btn.yellow:hover:not(:disabled) { background: #ca8a04; transform: translateY(-2px); }

    .btn.red { background: #ef4444; }
    .btn.red:hover:not(:disabled) { background: #dc2626; transform: translateY(-2px); }

    .input-field {
        padding: 0.5rem 0.75rem;
        font-size: 1rem;
        border-radius: 10px;
        background: rgba(0, 0, 0, 0.3);
        color: white;
        border: 1px solid rgba(255, 255, 255, 0.3);
        outline: none;
        text-align: center;
        width: 4.2rem;
        font-variant-numeric: tabular-nums;
    }

    .input-field.small {
        font-size: 0.95rem;
    }

    .input-field:focus {
        border-color: var(--primary-color, #6366f1);
    }
</style>
