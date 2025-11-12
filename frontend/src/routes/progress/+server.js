import { API_URL } from '$lib/config.js';

export async function POST({ request }) {
  try {
    const contentType = request.headers.get('content-type') || '';

    // ==========================
    // 1️⃣ Handle Multipart Upload (photo upload)
    // ==========================
    if (contentType.includes('multipart/form-data')) {
      const res = await fetch(`${API_URL}/upload`, {
        method: 'POST',
        body: request.body,
        duplex: 'half'
      });

      const data = await res.json().catch(() => ({}));
      return new Response(JSON.stringify(data), {
        status: res.status,
        headers: { 'Content-Type': 'application/json' }
      });
    }

    // ==========================
    // 2️⃣ Handle JSON Requests
    // ==========================
    const body = await request.json().catch(() => ({}));
    const { endpoint, data } = body || {};

    if (!endpoint) {
      return new Response(
        JSON.stringify({ success: false, message: 'Missing endpoint name' }),
        { status: 400, headers: { 'Content-Type': 'application/json' } }
      );
    }

    const backendRes = await fetch(`${API_URL}/${endpoint}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    });

    const contentTypeBackend = backendRes.headers.get('content-type') || '';
    let backendData;

    if (contentTypeBackend.includes('application/json')) {
      backendData = await backendRes.json();
    } else {
      const text = await backendRes.text();
      backendData = { success: backendRes.ok, message: text };
    }

    // ==========================
    // 3️⃣ Normalize for frontend expectations
    // ==========================

    // --- Case 1: getprogress returns an array ---
    if (endpoint === 'getprogress' && Array.isArray(backendData)) {
      const normalizedArray = backendData.map((entry) => ({
        id: entry.id || entry._id || entry.ID || crypto.randomUUID(),
        date: entry.date || '',
        weight: entry.weight || '',
        message: entry.message || '',
        photo: entry.photoBase64 || entry.photo || ''
      }));

      return new Response(JSON.stringify(normalizedArray), {
        status: 200,
        headers: { 'Content-Type': 'application/json' }
      });
    }

    // --- Case 2: addprogress, deleteprogress, or others ---
    let normalizedData = backendData?.data || backendData || null;

    if (normalizedData && typeof normalizedData === 'object') {
      // Normalize `_id` → `id`
      if (normalizedData._id && !normalizedData.id) {
        normalizedData.id = normalizedData._id;
      }

      // Guarantee `id` always exists for frontend reactivity
      if (!normalizedData.id && (endpoint === 'addprogress' || endpoint === 'getprogress')) {
        normalizedData.id = crypto.randomUUID();
      }
    }

    const normalized = {
      success: backendData.success ?? backendRes.ok,
      message: backendData.message || backendData.error || '',
      data: normalizedData
    };

    return new Response(JSON.stringify(normalized), {
      status: backendRes.status,
      headers: { 'Content-Type': 'application/json' }
    });

  } catch (err) {
    console.error('[Progress Proxy Error]', err);
    return new Response(
      JSON.stringify({
        success: false,
        message: 'Failed to contact backend',
        details: err.message
      }),
      {
        status: 502,
        headers: { 'Content-Type': 'application/json' }
      }
    );
  }
}
